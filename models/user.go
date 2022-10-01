package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName   string             `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName    string             `bson:"lastName,omitempty" json:"lastName,omitempty"`
	DateOfBirth time.Time          `bson:"dateOfBirth,omitempty" json:"dateOfBirth,omitempty"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password,omitempty"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempty"`
	Banner      string             `bson:"banner" json:"banner,omitempty"`
	Biography   string             `bson:"biography" json:"biography,omitempty"`
	Location    string             `bson:"location" json:"location,omitempty"`
	Website     string             `bson:"website" json:"website,omitempty"`
}
