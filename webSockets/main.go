package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

type Notification struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool)

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		clients[conn] = true

		go handleWebSocketConnection(conn)
	})

	r.POST("/notification", func(c *gin.Context) {
		var notification Notification
		err := c.BindJSON(&notification)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		sendNotification(notification)
	})

	r.Run(":8080")
}

func handleWebSocketConnection(conn *websocket.Conn) {
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			delete(clients, conn)
			break
		}
	}
}

func sendNotification(notification Notification) {
	fmt.Println(clients)
	for client := range clients {
		err := client.WriteJSON(notification)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}
