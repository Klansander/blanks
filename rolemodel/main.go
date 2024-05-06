package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func main() {
	// Создаем новый экземпляр Gin
	r := gin.Default()

	// Определяем маршруты и их обработчики
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.POST("/e", func(c *gin.Context) {
		var request struct {
			RegionID []uuid.UUID `json:"region_id"`
			IssueID  uuid.UUID   `json:"issue_id"`
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"echo": request})
	})

	// Запускаем сервер на порту 8080
	r.Run(":8080")
}
