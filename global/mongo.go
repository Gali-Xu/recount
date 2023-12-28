package global

import "go.mongodb.org/mongo-driver/mongo"

var (
	MongoClient *mongo.Client
	UserColl    *mongo.Collection
	ConsumeColl *mongo.Collection
	RoleColl    *mongo.Collection
)
