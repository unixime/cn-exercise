package model

import (
	"github.com/google/uuid"
)

type TYPE int

const (
	SENDING   TYPE = iota
	RECEIVING TYPE = iota
)

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
