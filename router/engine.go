package router

import (
	"github.com/gin-gonic/gin"
	"recount/controller"
	"recount/middleware"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()                         //初始化
	engine.Use(middleware.Cors())                   //跨域
	engine.POST("/register", controller.CreateUser) //注册
	engine.POST("/login", controller.Longin)        //登录
	engine.Use(middleware.JWTAuth())
	UserRouter(engine)
	ConsumeRouter(engine)
	return engine
}
