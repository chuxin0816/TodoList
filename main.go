package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/xxx", "./statics")
	r.LoadHTMLGlob("templates/*")
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/about.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})
	r.GET("/blog.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.html", nil)
	})
	r.GET("/client.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "client.html", nil)
	})
	r.GET("/contact.html",func(c *gin.Context){
		c.HTML(http.StatusOK,"contact.html",nil)
	})
	r.GET("/services.html",func(c *gin.Context){
		c.HTML(http.StatusOK,"services.html",nil)
	})
	r.Run("localhost:8080")
}
