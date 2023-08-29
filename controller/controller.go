package controller

import (
	"net/http"

	"github.com/chuxin0816/web-tutorial/models"
	"github.com/gin-gonic/gin"
)

func GetTodoList(ctx *gin.Context) {
	todoList, err := models.GetTodoList()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

func GetTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := models.GetTodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func PostTodo(ctx *gin.Context) {
	var todo models.Todo
	ctx.BindJSON(&todo)
	err := models.PostTodo(todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func PutTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	todo, err := models.GetTodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	ctx.BindJSON(&todo)
	err = models.PutTodo(todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	err := models.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
	}
}
