package client

import (
	"cn-exercise/internal/model"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateTransaction(t *testing.T) {

	customer := model.Customer{
		AccountID: 100,
		Name:      "John",
		IBAN:      1234567,
		Address:   "foo",
		Transactions: model.Transaction{
			Amount: 50,
			Type:   1,
		},
	}

	c := Client{
		url:    "https://vault.immudb.io",
		apiKey: "default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y",
	}

	e := c.RegisterTransaction("default", "default", &customer)

	require.Nil(t, e)

}

func TestLookForCustomerTransactions(t *testing.T) {
	c := Client{
		url:    "https://vault.immudb.io",
		apiKey: "default.AIkWyayo4M8uOBVUbce3zg.DyHDJbEg9chloDI6deZ2ldxERsi_z-fxifUqgkNuzsH5TZ3y",
	}

	e := c.LookForCustomerTransactions("default", "default", "John")

	require.Nil(t, e)
}
