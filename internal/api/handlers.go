package api

import (
	"cn-exercise/internal/client"
	"cn-exercise/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

// postAlbums adds an album from JSON received in the request body.
func PostTransaction(client *client.Client) gin.HandlerFunc {

	const ledger = "default"
	const collection = "transactions"

	fn := func(ctx *gin.Context) {
		var newTransaction model.Transaction

		if err := ctx.BindJSON(&newTransaction); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		response := client.RegisterTransaction(ledger, collection, &newTransaction)

		if response.Code != http.StatusOK {
			ctx.IndentedJSON(response.Code, gin.H{"message": response.Error()})
		}

		ctx.IndentedJSON(response.Code, gin.H{"message": "OK"})
	}

	return fn
}

func GetCustomerTransactions(httpClient *client.Client) gin.HandlerFunc {

	const ledger = "default"
	const collection = "transactions"

	fn := func(ctx *gin.Context) {

		if len(ctx.Request.URL.Query()) == 0 {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "need at least 1 param"})
			return
		}

		f := func(c *gin.Context) (func(l, c string, v string) (*client.Response, *model.SearchResponse), string, error) {

			if ctx.Request.URL.Query().Has("name") {
				return httpClient.GetTransactionByCustomerName, "name", nil
			}

			if ctx.Request.URL.Query().Has("uuid") {
				return httpClient.GetTransactionByCustomerUUID, "uuid", nil
			}

			//ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not supported query parameter"})

			return nil, "", errors.New("not supported query parameter")

		}

		hdl, key, e := f(ctx)
		if e != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "not supported query parameter"})

			return
		}

		r, data := hdl(ledger, collection, ctx.Query(key))
		ctx.IndentedJSON(r.Code, data)

		return
		//customerName := ctx.Query("name")
		//if customerName != "" {
		//	r, data := client.GetTransactionByCustomerName(ledger, collection, customerName)
		//
		//	ctx.IndentedJSON(r.Code, data)
		//
		//	return
		//
		//}
		//
		//customerUUID := ctx.Query("id")
		//if customerUUID != "" {
		//	r, data := client.GetTransactionByCustomerUUID(ledger, collection, customerUUID)
		//
		//	ctx.IndentedJSON(r.Code, data)
		//
		//	return
		//
		//}

	}

	return fn
}
