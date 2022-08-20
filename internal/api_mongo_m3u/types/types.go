package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Track struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	URI  string             `json:"uri" bson:"uri"`
}

type PostTrack struct {
	Name string `json:"name" bson:"name"`
	URI  string `json:"uri" bson:"uri"`
}

type DeleteTrack struct {
	Id primitive.ObjectID `json:"_id" bson:"_id"`
}
