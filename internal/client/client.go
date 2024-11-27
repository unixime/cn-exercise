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

func (c *ImmuDBClientIfImpl) IsConnected() bool {
	return c.client.IsConnected()
}

func (c *ImmuDBClientIfImpl) CreateDatabase(name string) error {
	_, err := c.client.CreateDatabaseV2(context.Background(), name, nil)

	//if s.AlreadyExisted {
	//	return nil
	//}
	//if errors.Code(err) == errors.CodInternalError {
	//	return nil
	//}

	return err
}

func (c *ImmuDBClientIfImpl) CreateTable(tableName string, tableDef string) error {

	params := map[string]interface{}{"name": tableName}
	_, err := c.client.SQLExec(context.Background(), tableDef, params)

	return err
}
