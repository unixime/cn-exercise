package handlers

import (
	"cn-exercise/internal/client"
	"cn-exercise/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func AddTransactionHandler(client client.ImmuDBClientIf) gin.HandlerFunc {

	fn := func(ctx *gin.Context) {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
		data := model.Payload{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			panic(err)
		}

		fmt.Println(client.IsConnected())
	}

	return fn
}
