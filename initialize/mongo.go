package initialize

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"recount/global"
)

func MongoInit() {
	if global.MongoClient == nil {
		global.MongoClient = getMongoClient("mongodb://admin:zsx2003@101.43.217.23:27017")
	}
	recount := global.MongoClient.Database("recount")
	{
		global.UserColl = recount.Collection("User")
		global.ConsumeColl = recount.Collection("Consume")
	}
}

func getMongoClient(url string) *mongo.Client {
	//设置客户端选项
	clientOptions := options.Client().ApplyURI(url)
	//连接MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
