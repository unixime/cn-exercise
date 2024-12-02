package cmd

import (
	"cn-exercise/internal/api"
	"cn-exercise/internal/client"

	_ "cn-exercise/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           CN API
// @version         1.0
// @description     Simple server that allows adding and getting transaction from immudb vault.

// @contact.name   Emanuele Piccinelli
// @contact.email  emanuele.piccinelli@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func StartServer(client client.Client) error {

	server := gin.Default()

	server.POST("/transaction", api.PostTransaction(&client))

	server.GET("/transactions", api.GetCustomerTransactions(&client))

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
