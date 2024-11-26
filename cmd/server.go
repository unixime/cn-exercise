package cmd

import (
	"cn-exercise/internal/client"

	"github.com/gin-gonic/gin"
)

func StartServer() error {

	c := client.NewImmuDBClient()

	c.Login("127.0.0.1", 3322, "immudb", "immudb")

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return server.Run() // listen and serve on 0.0.0.0:8080
}
