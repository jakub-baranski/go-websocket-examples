package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// Default cors settings - allow all origins
	r.Use(cors.Default())

	// HTML Handlers

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"simpleEchoWebsocket": "/echo/",
			"base64Cat":           "/cat/",
			"persistentChat":      "/chat/",
		})
	})

	r.GET("/test/:id", func(c *gin.Context) {
		param, _ := c.Params.Get("id")
		c.JSON(200, gin.H{
			"test": param,
		})
	})

	// Websocket handlers

	r.GET("/echo/", func(c *gin.Context) {
		EchoHandler(c.Writer, c.Request)
	})

	r.GET("/cat/", func(c *gin.Context) {
		CatHandler(c.Writer, c.Request)
	})

	r.GET("/chat/", func(c *gin.Context) {
		ChatHandler(c.Writer, c.Request)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
