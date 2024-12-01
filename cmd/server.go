package cmd

import (
	"cn-exercise/internal/api"
	"cn-exercise/internal/client"

	_ "cn-exercise/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func StartServer(client client.Client) error {

	server := gin.Default()

	server.POST("/transaction", api.PostTransaction(&client))

	server.GET("/transactions", api.GetCustomerTransactions(&client))

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return server.Run(":8080") // listen and serve on 0.0.0.0:8080
}
