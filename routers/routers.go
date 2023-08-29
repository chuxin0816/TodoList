package routers

import (
	"net/http"

	"github.com/chuxin0816/web-tutorial/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")
	r.Static("/static", "static")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	v1Group := r.Group("/v1/todo")
	{
		v1Group.GET("/", controller.GetTodoList)
		v1Group.GET("/:id", controller.GetTodo)
		v1Group.POST("/", controller.PostTodo)
		v1Group.PUT("/:id", controller.PutTodo)
		v1Group.DELETE("/:id", controller.DeleteTodo)
	}
	return r
}
