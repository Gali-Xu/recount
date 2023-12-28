package service

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"recount/global"
	"recount/model"
	"recount/utils"
	"strconv"
)

func CreateConsume(consume model.Consume) utils.Response {
	res, err := global.ConsumeColl.InsertOne(context.TODO(), consume)
	if err != nil {
		fmt.Println(consume.UserId)
		return utils.ErrorMess("添加失败", err)
	}
	return utils.SuccessMess("添加成功", res)
}

func DeleteConsume(id primitive.ObjectID) utils.Response {
	if err := global.ConsumeColl.FindOne(context.TODO(), bson.M{"_id": id}).Decode(bson.M{}); err != nil {
		return utils.ErrorMess("数据库中无此消费记录", err)
	}
	res, err := global.ConsumeColl.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return utils.ErrorMess("删除失败", err)
	}
	return utils.SuccessMess("删除成功", res.DeletedCount)
}

func UpdateConsume(consume model.Consume) utils.Response {
	if consume.Id.IsZero() {
		return utils.ErrorMess("数据库中无此消费记录", nil)
	}
	res, err := global.ConsumeColl.UpdateByID(context.TODO(), bson.M{"_id": consume.Id}, bson.M{"$set": bson.M{"cost": consume.Cost,
		"category": consume.Category}})
	if err != nil {
		return utils.ErrorMess("更新失败", err)
	}
	return utils.SuccessMess("更新成功", res)
}

func GetConsume(userId primitive.ObjectID, cost float64, costTime string, category, currPage, pageSize string) utils.Response {
	curr, _ := strconv.ParseInt(currPage, 10, 64)
	page, _ := strconv.ParseInt(pageSize, 10, 64)
	skip := page * (curr - 1)
	var filter bson.M
	if costTime == "" {
		filter = bson.M{
			"userId":   userId,
			"category": bson.M{"$regex": category},
		}
	} else {
		filter = bson.M{
			"userId":   userId,
			"category": bson.M{"$regex": category},
			"costTime": bson.M{"$regex": costTime},
		}
	}
	opts := options.FindOptions{
		Limit: &page,
		Skip:  &skip,
		Sort:  bson.M{"_id": -1},
	}
	res, err := global.ConsumeColl.Find(context.TODO(), filter, &opts)
	if err != nil {
		return utils.ErrorMess("查询失败", err)
	}
	fmt.Println(res)
	var consumeDB []model.Consume
	if err := res.All(context.TODO(), &consumeDB); err != nil {
		return utils.ErrorMess("归档失败", err)
	}
	fmt.Println(consumeDB)
	var sum = 0.0
	for i, _ := range consumeDB {
		sum += consumeDB[i].Cost
	}
	data := map[string]interface{}{
		"Data": consumeDB,
		"sum":  sum,
	}
	return utils.SuccessMess("查询成功", data)
}
