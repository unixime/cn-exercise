package client

import (
	"cn-exercise/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateTransaction(t *testing.T) {

	const ledger = "default"
	const collection = "transactions"

	tx := model.Transaction{
		AccountID: "100",
		Name:      "John Blake",
		IBAN:      "IT32C0300203280141759145451",
		Address:   "foo",
		Amount:    50,
		Type:      1,
	}

	c := Client{
		url:    "https://vault.immudb.io",
		apiKey: "default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y",
	}

	e := c.RegisterTransaction(ledger, collection, &tx)

	require.Nil(t, e)

}

func TestGetTransactionByCustomerName(t *testing.T) {

	const ledger = "default"
	const collection = "transactions"

	c := Client{
		url:    "https://vault.immudb.io",
		apiKey: "default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y",
	}

	e := c.GetTransactionByCustomerName(ledger, collection, "John Blake")

	require.Nil(t, e)
}

func TestGetTransactionByCustomerUUID(t *testing.T) {

	const ledger = "default"
	const collection = "transactions"

	c := Client{
		url:    "https://vault.immudb.io",
		apiKey: "default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y",
	}

	e := c.GetTransactionByCustomerUUID(ledger, collection, "100")

	require.Nil(t, e)
}

func TestCollectionExists(t *testing.T) {
	const ledger = "default"
	const collection = "transactions"

	c := Client{
		url:    "https://vault.immudb.io",
		apiKey: "default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y",
	}

	_, e := c.CollectionExists(ledger, collection)

	require.Nil(t, e)
}
