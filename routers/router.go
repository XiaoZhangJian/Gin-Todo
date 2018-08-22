package routers

import (
	"Gin-Todo/middleware/jwt"
	"Gin-Todo/pkg/setting"
	"Gin-Todo/routers/api"
	"Gin-Todo/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.Default()
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/login", api.Login)
	r.POST("/reg", api.Register)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		// 用户操作
		// apiv1.PUT("/users/:id", api.EditUser)

		//获取任务
		apiv1.GET("/todos", v1.GetTodos)
		apiv1.GET("/todos/:id", v1.GetTodo)
		//新建任务
		apiv1.POST("/todos", v1.AddTodo)
		//更新指定任务
		apiv1.PUT("/todos/:id", v1.EditTodo)
		//删除指定任务
		apiv1.DELETE("/todos/:id", v1.DelTodo)
	}

	return r
}
