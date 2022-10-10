package database

import (
	"context"
	"fmt"
	"time"

	"github.com/godoquin/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsersBD(ID string, page int64, search string, types string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")
	var results []*models.User
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 5)
	findOptions.SetLimit(5)
	query := bson.M{
		"firstName": bson.M{"$regex": `(?i)` + search},
	}
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var finded, include bool
	for cur.Next(ctx) {
		var s models.User
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()
		include = false
		finded, _ = GetRelationBD(r)
		if types == "new" && !finded {
			include = true
		}
		if types == "follow" && finded {
			include = true
		}
		if r.UserRelationID == ID {
			include = false
		}
		if include {
			s.Password = ""
			s.Biography = ""
			s.Website = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""
			results = append(results, &s)
		}

	}
	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
