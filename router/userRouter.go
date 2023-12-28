package router

import (
	"github.com/gin-gonic/gin"
	"recount/controller"
)

func UserRouter(engine *gin.Engine) {
	user := engine.Group("/user")
	{
		user.POST("/create", controller.CreateUser)
		user.DELETE("/delete", controller.DeleteUser)
		user.PUT("/update", controller.UpdateUser)
	}
}
