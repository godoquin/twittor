package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTweetBD(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	ObjID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":    ObjID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(ctx, condition)
	return err
}
