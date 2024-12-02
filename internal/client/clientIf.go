package client

import "cn-exercise/internal/model"

type ClientIf interface {
	RegisterTransaction(ledger string, collection string, transaction *model.Transaction) *Response
	GetTransactionByCustomerName(ledger string, collection string, name string) (*Response, *model.SearchResponse)
	GetTransactionByCustomerUUID(ledger string, collection string, uuid string) (*Response, *model.SearchResponse)
}
