package database

import (
	"context"
	"fmt"
	"time"

	"github.com/godoquin/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")
	var profile models.User
	objtID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objtID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Registro no econtrado " + err.Error())
		return profile, err
	}
	return profile, nil
}
