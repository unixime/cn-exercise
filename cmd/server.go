package cmd

import (
	"cn-exercise/internal/api"
	"cn-exercise/internal/client"
	"cn-exercise/internal/handlers"
	"github.com/gin-gonic/gin"
)

func StartServer(client client.Client) error {

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//server.POST("/foo", handlers.AddTransactionHandler(c))
	server.GET("/customer", handlers.CustomerTransactionHandler())

	//server.POST("/transaction", api.PostTransaction)
	server.POST("/transaction", api.PostTransaction(&client))

	server.GET("/transactions", api.GetCustomerTransactions(&client))

	return server.Run( /*:9888*/ ) // listen and serve on 0.0.0.0:8080
}
