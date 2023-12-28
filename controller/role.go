package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"recount/model"
	"recount/service"
)

func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, service.CreateRole(role))
}

func DeleteRole(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Query("_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, service.DeleteRole(id))
}

func UpdateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, service.UpdateRole(role))
}

func GetRole(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Query("_id"))
	curr := c.DefaultQuery("currpage", "1")
	page := c.DefaultQuery("pagesize", "0")
	rolename := c.Query("rolename")
	c.JSON(http.StatusOK, service.GetRole(id, rolename, curr, page))
}
