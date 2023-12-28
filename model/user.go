package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	Account   string             `json:"account" bson:"account"`
	AvatarUrl string             `json:"avatarUrl" bson:"avatarUrl"`
	Salt      string             `json:"salt" bson:"salt"`
}
