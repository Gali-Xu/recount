package router

import (
	"github.com/gin-gonic/gin"
	"recount/controller"
)

func ConsumeRouter(engine *gin.Engine) {
	consume := engine.Group("/consume")
	{
		consume.POST("/create", controller.CreateConsume)   //新增消费记录
		consume.DELETE("/delete", controller.DeleteComsume) //删除一条消费记录
		consume.PUT("/update", controller.UpdateConsume)    //修改消费记录
		consume.GET("/get", controller.GetConsume)          //根据条件查询消费记录
	}
}
