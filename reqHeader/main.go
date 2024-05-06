package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	// Force log's color
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {

		for k, v := range c.Request.Header {
			fmt.Println(k, v)
		}
		c.String(200, "pong")

	})

	router.Run(":8080")
}
