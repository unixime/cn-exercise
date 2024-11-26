package client

import (
	"context"
	"log"

	immudb "github.com/codenotary/immudb/pkg/client"
)

type ImmuDBClientIfImpl struct {
	client immudb.ImmuClient
}

func NewImmuDBClient() *ImmuDBClientIfImpl {
	return &ImmuDBClientIfImpl{}
}

func (c *ImmuDBClientIfImpl) Login(host string, port int, user string, password string) error {
	opts := immudb.DefaultOptions().
		WithAddress(host).
		WithPort(port)

	c.client = immudb.NewClient().WithOptions(opts)
	err := c.client.OpenSession(
		context.TODO(),
		[]byte(user),
		[]byte(password),
		"defaultdb",
	)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (c *ImmuDBClientIfImpl) Logout() error {
	return c.client.CloseSession(context.TODO())
}
