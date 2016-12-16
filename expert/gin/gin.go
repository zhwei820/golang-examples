package main

import "github.com/gin-gonic/gin"
import "time"

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		c.String(200, "hello world")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.POST("/submit", func(c *gin.Context) {
		c.String(401, "not authorized")
	})
	router.PUT("/error", func(c *gin.Context) {
		c.String(500, "and error hapenned :(")
	})
	router.Run(":23333")
}
