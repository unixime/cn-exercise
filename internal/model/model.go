package model

import (
	"encoding/json"
	"github.com/google/uuid"
)

type TYPE int

const (
	SENDING   TYPE = iota
	RECEIVING TYPE = iota
)

type Collection interface {
	Name() string
	AsJSON() []byte
}

type Customers struct {
}

func (c *Customers) Name() string {
	const name = "customers"

	return name
}

func (c *Customers) AsJSON() []byte {
	const data = `
	{
  		"idFieldName": "customers",
  		"fields": [
			{
			  "name": "string",
			  "type": "STRING"
			},
			{
				name: "accountid",
				type: "STRING",
			},
			{
				"name": "address",
				"type": "STRING",
			},
			{
				"name": "iban",
				"type": "STRING",	
			},
			{
				"name": "balance",
				"type": "DOUBLE",
			}
	  	],
		"indexes": [
			{
			  "fields": [
				"name",
				"iban"
			  ],
			  "isUnique": true
			}
	  	]
	}	`

	return []byte(data)
}

type Transactions struct {
	name string
}

func (t *Transactions) Name() string {
	return t.name
}

func (t *Transactions) AsJSON() []byte {
	const def = `
	{
	  "fields": [
	  {
		"name": "uuid",
		"type": "STRING"
	  },
	  {
		"name": "name",
		"type": "STRING"
	  },
	  {
		"name": "iban",
		"type": "STRING"
	  },
	  {
		"name": "address",
		"type": "STRING"
	  },
	  {
		"name": "amount",
		"type": "DOUBLE"
	  },
	  {
		"name": "type",
		"type": "INTEGER"
	  }
	  ],
	  "indexes": [
	  {
		"fields": [
		  "uuid"
		],
		"isUnique": false
	  },
	  {
		"fields": [
		  "name"
		],
		"isUnique": false
	  },
	  {
		"fields": [
		  "name",
		  "type"
		],
		"isUnique": false
	  }
	  ]
	}`
	return []byte(def)

}

//func (c *Transactions) AsJSON() []byte {
//	const data = `
//	{
//  		"idFieldName": "transactions",
//  		"fields": [
//			{
//				name: "accountid",
//				type: "STRING",
//			},
//			{
//			  "name": "amount",
//			  "type": "DOUBLE"
//	  	],
//		"indexes": [
//			{
//			  "fields": [
//				"accountid"
//			  ],
//			  "isUnique": false
//			}
//	  	]
//	}	`
//
//	return []byte(data)
//}

type Payload struct {
	AccountID   uuid.UUID
	Name        string  `json:"name"`
	BankAccount int     `json:"bank_account"`
	Address     string  `json:"address"`
	Amount      float64 `json:"amount"`
	Type        TYPE    `json:"type"`
}

func NewPayload(name string) *Payload {

	return &Payload{
		AccountID:   uuid.New(),
		Name:        name,
		BankAccount: 0,
		Address:     "",
		Amount:      0,
		Type:        0,
	}

}

type Transaction struct {
	AccountID string  `json:"uuid"`
	Name      string  `json:"name"`
	IBAN      string  `json:"iban"`
	Address   string  `json:"address,omitempty"`
	Amount    float64 `json:"amount"`
	Type      TYPE    `json:"type"`
}

func (c *Transaction) AsJSON() ([]byte, error) {

	data, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	return data, err
}

type Revision struct {
	TransactionId string      `json:"transactionId,omitempty"`
	Revision      string      `json:"revision,omitempty"`
	Transaction   Transaction `json:"document,omitempty"`
}

type SearchResponse struct {
	SearchId  string     `json:"searchId,omitempty"`
	KeepOpen  bool       `json:"keepOpen,omitempty"`
	Revisions []Revision `json:"revisions,omitempty"`
	Page      int        `json:"page"`
	PerPage   int        `json:"perPage"`
}

type Customer struct {
	AccountID    int    `json:"accountid"`
	Name         string `json:"name"`
	IBAN         int    `json:"iban"`
	Address      string `json:"address,omitempty"`
	Transactions Transaction
}

func (c *Customer) AsJSON() ([]byte, error) {

	data, err := json.Marshal(c)

	if err != nil {
		return nil, err
	}

	return data, err
}
