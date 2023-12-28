package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Rolename  string             `json:"rolename" bson:"rolename"`
	Authority string             `json:"authority" bson:"authority"`
}
