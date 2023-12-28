package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Consume struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId   primitive.ObjectID `json:"userId" bson:"userId"`
	Cost     float64            `json:"cost" bson:"cost"`
	CostTime string             `json:"costTime" bson:"costTime"`
	Category string             `json:"category" bson:"category"`
}
