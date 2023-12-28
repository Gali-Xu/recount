package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"recount/global"
	"recount/middleware"
	"recount/model"
	"recount/utils"
	"strconv"
	"time"
)

func CreateUser(user model.User) utils.Response {
	if err := global.UserColl.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(bson.M{}); err == nil {
		return utils.ErrorMess("账号已存在", err)
	}
	rand.Seed(time.Now().Unix()) //根据时间戳生成种子
	//生成盐
	salt := strconv.FormatInt(rand.Int63(), 10)
	//密码加盐加密
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return utils.ErrorMess("密码加密失败", err)
	}
	user.Password, user.Salt = string(encryptedPass), salt
	res, err := global.UserColl.InsertOne(context.TODO(), user)
	if err != nil {
		return utils.ErrorMess("添加失败", err)
	}
	return utils.SuccessMess("添加用户成功", res)
}

func Longin(user model.User) utils.Response {
	var DBUser model.User
	//校验账号
	err := global.UserColl.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&DBUser)
	if err != nil {
		return utils.ErrorMess("账号错误", err.Error())
	}
	//校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(user.Password+DBUser.Salt)); err != nil {
		return utils.ErrorMess("密码错误", err.Error())
	}
	//生成token
	token, err := middleware.CreateToken(DBUser)
	if err != nil {
		return utils.ErrorMess("生成token失败", err.Error())
	}
	res := map[string]interface{}{
		"_id":       DBUser.Id,
		"account":   DBUser.Account,
		"name":      DBUser.Username,
		"password":  DBUser.Password,
		"avatarUrl": DBUser.AvatarUrl,
		"token":     token,
	}
	return utils.SuccessMess("登陆成功", res)
}

func DeleteUser(id primitive.ObjectID) utils.Response {
	res, err := global.UserColl.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return utils.ErrorMess("删除失败", err)
	}
	if res.DeletedCount == 0 {
		return utils.ErrorMess("删除不存在", err)
	}
	return utils.SuccessMess("删除成功", res)
}

func UpdateUser(user model.User) utils.Response {
	if user.Id.IsZero() {
		return utils.ErrorMess("更新用户不存在", nil)
	}
	res, err := global.UserColl.UpdateByID(context.TODO(), bson.M{"_id": user.Id}, bson.M{"$set": bson.M{"password": user.Password,
		"account": user.Account, "avatarUrl": user.AvatarUrl}})
	if err != nil {
		return utils.ErrorMess("更新失败", err)
	}
	return utils.SuccessMess("更新成功", res.UpsertedID)
}
