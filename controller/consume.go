package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"recount/model"
	"recount/service"
	"strconv"
	"time"
)

func CreateConsume(c *gin.Context) {
	var consume model.Consume
	if err := c.Bind(&consume); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	consume.CostTime = time.Now().Format("2006-01-02 15:04:05")
	c.JSON(http.StatusOK, service.CreateConsume(consume))
}

func DeleteComsume(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, service.DeleteConsume(id))
}

func UpdateConsume(c *gin.Context) {
	var consume model.Consume
	if err := c.Bind(&consume); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
	}
	c.JSON(http.StatusOK, service.UpdateConsume(consume))
}

func GetConsume(c *gin.Context) {
	userId, _ := primitive.ObjectIDFromHex(c.Query("userId"))
	cost, _ := strconv.ParseFloat(c.Query("cost"), 64)
	costTime := c.Query("costTime")
	category := c.DefaultQuery("category", "")
	currPage := c.DefaultQuery("currPage", "1")
	pageSize := c.DefaultQuery("pageSize", "0")
	fmt.Println(userId)
	c.JSON(http.StatusOK, service.GetConsume(userId, cost, costTime, category, currPage, pageSize))
}
