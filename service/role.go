package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"recount/global"
	"recount/model"
	"recount/utils"
	"strconv"
)

func CreateRole(role model.Role) utils.Response {
	if err := global.RoleColl.FindOne(context.TODO(), bson.M{"rolename": role.Rolename}).Decode(bson.M{}); err != nil {
		utils.ErrorMess("角色名重复", err)
	}
	res, err := global.RoleColl.InsertOne(context.TODO(), role)
	if err != nil {
		utils.ErrorMess("创建失败", err)
	}
	return utils.SuccessMess("创建成功", res)
}

func DeleteRole(id primitive.ObjectID) utils.Response {
	var deleteRole bson.M
	err := global.RoleColl.FindOneAndDelete(context.TODO(), bson.M{"_id": id}).Decode(&deleteRole)
	if err != nil {
		utils.ErrorMess("无法找到或删除失败", err)
	}
	return utils.SuccessMess("删除成功", deleteRole)
}

func UpdateRole(role model.Role) utils.Response {
	if role.Id.IsZero() {
		utils.ErrorMess("该角色id不存在", nil)
	}
	res, err := global.RoleColl.UpdateByID(context.TODO(), bson.M{"_id": role.Id}, bson.M{"$set": role})
	if err != nil {
		return utils.ErrorMess("更新失败", err)
	}
	return utils.SuccessMess("更新成功", res)
}

func GetRole(id primitive.ObjectID, rolename, currpage, pagesize string) utils.Response {
	curr, _ := strconv.ParseInt(currpage, 10, 64)
	page, _ := strconv.ParseInt(pagesize, 10, 64)

	skip := (curr - 1) * page
	filter := bson.M{
		"rolename": bson.M{"$regex": rolename},
	}
	opts := options.FindOptions{
		Limit: &page,
		Skip:  &skip,
		Sort:  bson.M{"_id": -1},
	}
	cur, err := global.RoleColl.Find(context.TODO(), filter, &opts)
	if err != nil {
		return utils.ErrorMess("查询失败", err)
	}
	var resultDB []model.Role
	if err := cur.All(context.TODO(), &resultDB); err != nil {
		return utils.ErrorMess("归档失败", err)
	}
	return utils.SuccessMess("查询成功", resultDB)
}
