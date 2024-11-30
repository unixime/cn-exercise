package cmd

import (
	"cn-exercise/internal/handlers"

	"github.com/gin-gonic/gin"
)

func StartServer() error {

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//server.POST("/foo", handlers.AddTransactionHandler(c))
	server.GET("/customer", handlers.CustomerTransactionHandler())
	return server.Run( /*:9888*/ ) // listen and serve on 0.0.0.0:8080
}
