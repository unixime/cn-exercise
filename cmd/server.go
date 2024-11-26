package cmd

import (
	"cn-exercise/internal/client"
	"cn-exercise/internal/handlers"
	"github.com/gin-gonic/gin"
)

func StartServer() error {

	c := client.NewImmuDBClient()

	c.Login("127.0.0.1", 3322, "immudb", "immudb")
	defer c.Logout()

	if e := c.CreateDatabase("customer_tx"); e != nil {
		return e
	}

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.POST("/foo", handlers.AddTransactionHandler(c))
	return server.Run( /*:9888*/) // listen and serve on 0.0.0.0:8080
}
