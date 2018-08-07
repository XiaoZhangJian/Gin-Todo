package routers

import (
	"Gin-Todo/middleware/jwt"
	"Gin-Todo/pkg/setting"
	"Gin-Todo/routers/api"
	"Gin-Todo/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取任务
		apiv1.GET("/tasks", v1.GetTasks)
		//新建任务
		apiv1.POST("/tasks", v1.AddTask)
		//更新指定任务
		apiv1.PUT("/tasks/:id", v1.EditTask)
		//删除指定任务
		apiv1.DELETE("/tasks/:id", v1.DeleteTask)
	}

	return r
}
