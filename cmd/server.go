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

	//if e := c.CreateDatabase("customer_tx"); e != nil {
	//	return e
	//}

	table := `CREATE TABLE IF NOT EXISTS transactions (
		accountid          UUID,
		name     VARCHAR NOT NULL,
		bank_account       INT NOT NULL,
		address  VARCHAR NOT NULL,
		amount FLOAT NOT NULL,
		tx_type VARCHAR NOT NULL,
		PRIMARY KEY (accountid)
	);`

	table = "CREATE TABLE IF NOT EXISTS transactions (" +
		"accountid  VARCHAR[256] NOT NULL," +
		"name     VARCHAR NOT NULL," +
		"bank_account       INTEGER NOT NULL," +
		"address  VARCHAR NOT NULL," +
		"amount FLOAT NOT NULL," +
		"tx_type VARCHAR NOT NULL," +
		"PRIMARY KEY (accountid));"

	c.CreateTable("table", table)

	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.POST("/foo", handlers.AddTransactionHandler(c))
	return server.Run( /*:9888*/) // listen and serve on 0.0.0.0:8080
}
