package database

import (
	"context"
	"log"
	"time"

	"github.com/godoquin/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadTweetBD(ID string, page int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var results []*models.ReturnTweets
	condition := bson.M{
		"userid": ID,
	}
	options := options.Find()
	options.SetLimit(5)
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	options.SetSkip((page - 1) * 5)

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}
	for cursor.Next(context.TODO()) {
		var reg models.ReturnTweets
		err := cursor.Decode(&reg)
		if err != nil {
			return results, false
		}
		results = append(results, &reg)
	}
	return results, true
}
