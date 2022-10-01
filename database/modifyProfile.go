package database

import (
	"context"
	"time"

	"github.com/godoquin/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyProfile(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")
	reg := make(map[string]interface{})
	if len(u.FirstName) > 0 {
		reg["firstName"] = u.FirstName
	}
	if len(u.LastName) > 0 {
		reg["lastName"] = u.LastName
	}
	reg["dateOfBirth"] = u.DateOfBirth
	if len(u.Avatar) > 0 {
		reg["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		reg["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		reg["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		reg["location"] = u.Location
	}
	if len(u.Website) > 0 {
		reg["website"] = u.Website
	}
	updtString := bson.M{
		"$set": reg,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}
