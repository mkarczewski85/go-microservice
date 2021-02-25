package main

import (
	"words-microservice/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong!")
	})

	r.POST("/test", controller.SearchForWordsHandler)

	r.Run(":3000")
}
