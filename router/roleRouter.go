package router

import "github.com/gin-gonic/gin"

func RoleRouter(engine *gin.Engine) {
	role := engine.Group("/role")
	{
		role.POST("/create")
	}
}
