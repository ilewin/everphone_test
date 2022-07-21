package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Interests []string           `bson:"interests" json:"interests"`
	Gifts     []Gift             `bson:"gifts,omitempty" json:"gifts,omitempty"`
}

type Gift struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `bson:"name" json:"name"`
	Categories []string           `bson:"categories" json:"categories"`
}
