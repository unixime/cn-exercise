package cmd

import (
	"cn-exercise/internal/handlers"

	"github.com/gin-gonic/gin"
)

func StartServer() error {

	//c := client.NewImmuDBClient()
	//
	//c.Login("127.0.0.1", 3322, "immudb", "immudb")
	//defer c.Logout()

	//if e := c.CreateDatabase("customer_tx"); e != nil {
	//	return e
	//}

	//table := `CREATE TABLE IF NOT EXISTS transactions (
	//	accountid          UUID,
	//	name     VARCHAR NOT NULL,
	//	bank_account       INT NOT NULL,
	//	address  VARCHAR NOT NULL,
	//	amount FLOAT NOT NULL,
	//	tx_type VARCHAR NOT NULL,
	//	PRIMARY KEY (accountid)
	//);`
	//
	//table = "CREATE TABLE IF NOT EXISTS transactions2 (" +
	//	"accountid  UUID," +
	//	"name     VARCHAR[64] NOT NULL," +
	//	"iban       INTEGER NOT NULL," +
	//	"address  VARCHAR[256] NOT NULL," +
	//	"amount FLOAT NOT NULL," +
	//	"tx_type INTEGER NOT NULL," +
	//	"PRIMARY KEY (accountid));" +
	//	"CREATE INDEX IF NOT EXISTS ON transactions2(name, iban);" +
	//	"CREATE INDEX IF NOT EXISTS ON transactions2(iban);" +
	//	"CREATE INDEX IF NOT EXISTS ON transactions2(name);"
	//
	//c.CreateTable("table", table)

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
