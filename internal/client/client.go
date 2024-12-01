package client

import (
	"bytes"
	"cn-exercise/internal/model"
	"cn-exercise/internal/query"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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

func (c *ImmuDBClientIfImpl) Insert(tableName string, payload model.Payload) error {
	tx, err := c.client.NewTx(context.Background())
	if err != nil {
		return err
	}

	eTx := tx.SQLExec(
		context.Background(),
		`INSERT INTO @table VALUES(accountid, name, bank_account, address, amount, tx_type);`,
		map[string]interface{}{
			"accountid":    payload.AccountID,
			"name":         payload.Name,
			"bank_account": payload.BankAccount,
			"address":      payload.Address,
			"amount":       payload.Amount,
			"type":         payload.Type},
	)

	if eTx != nil {
		return eTx
	}

	return nil
}

type Client struct {
	URL    string
	ApiKey string
}

type Response struct {
	Code int
	Msg  string
}

func (r *Response) Error() string {
	return fmt.Sprintf("%d: %s", r.Code, r.Msg)
}

func NewClient(apiKey string, url string) *Client {

	return &Client{ApiKey: apiKey, URL: url}
}

func (c *Client) Login(host string, port int, user string, password string) error {

	return nil
}

func (c *Client) CollectionExists(ledger, collection string) (bool, error) {
	const path = "/ics/api/v1/ledger"
	reqURL := fmt.Sprintf("%s/%s/%s/collection/%s", c.URL, path, ledger, collection)
	resp, err := http.Get(reqURL)
	if err != nil {
		return false, err
	}

	switch resp.StatusCode {
	case 200:
		return false, nil
	}

	fmt.Println(resp.Body)
	return true, nil
}

func (c *Client) NewCollection(ledger string, collection model.Collection) error {

	const path = "/ics/api/v1/ledger"
	const contentType = "application/json"

	reqURL := fmt.Sprintf("%s/%s/%s/collection/%s", c.URL, path, ledger, collection.Name())

	//https: //vault.immudb.io/ics/api/v1/ledger/{ledger}/collection/{collection}
	fmt.Println(collection.AsJSON())
	req, err := http.NewRequest("PUT", reqURL, bytes.NewBuffer(collection.AsJSON()))

	if err != nil {
		return err
	}

	req.Header.Set("accept", ": */*")
	req.Header.Set("X-API-Key", c.ApiKey)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (c *Client) RegisterTransaction(ledger string, collection string, transaction *model.Transaction) *Response {
	const path = "ics/api/v1/ledger"
	const contentType = "application/json"

	reqURL := fmt.Sprintf("%s/%s/%s/collection/%s/document", c.URL, path, ledger, collection)

	payload, e := transaction.AsJSON()

	if e != nil {
		return &Response{
			Code: http.StatusInternalServerError,
			Msg:  e.Error(),
		}
	}

	req, err := http.NewRequest("PUT", reqURL, bytes.NewBuffer(payload))

	if err != nil {
		return &Response{
			Code: http.StatusInternalServerError,
			Msg:  e.Error(),
		}
	}

	req.Header.Set("accept", ": */*")
	req.Header.Set("X-API-Key", c.ApiKey)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	return &Response{
		Code: resp.StatusCode,
		Msg:  resp.Status,
	}
}

func (c *Client) GetTransactionByCustomerName(ledger string, collection string, name string) (*Response, *model.SearchResponse) {

	const path = "ics/api/v1/ledger"
	const contentType = "application/json"

	reqURL := fmt.Sprintf("%s/%s/%s/collection/%s/documents/search", c.URL, path, ledger, collection)

	qry := query.Query{
		Expressions: []query.Expression{
			{
				[]query.Constraint{
					{
						Field:    "name",
						Operator: query.EQUAL,
						Value:    name,
					},
				},
			},
		},
		OrderBy: []query.OrderConstraint{
			{
				Field: "_vault_md.ts",
				Order: query.DESC,
			},
		},
		Limit: 0,
	}

	search := query.Search{
		KeepOpen: true,
		Query:    qry,
		Page:     1,
		PerPage:  100,
	}

	data, err := json.Marshal(search)

	if err != nil {
		return &Response{Code: http.StatusInternalServerError, Msg: err.Error()}, nil
	}

	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(data))

	if err != nil {
		return &Response{Code: http.StatusInternalServerError, Msg: err.Error()}, nil
	}

	req.Header.Set("accept", ": */*")
	req.Header.Set("X-API-Key", c.ApiKey)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return &Response{Code: http.StatusInternalServerError, Msg: err.Error()}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Response{Code: http.StatusInternalServerError, Msg: err.Error()}, nil
	}
	// close response body
	var docs model.SearchResponse
	if err := json.Unmarshal(body, &docs); err != nil {
		return &Response{Code: http.StatusInternalServerError, Msg: err.Error()}, nil
	}

	return &Response{
		Code: resp.StatusCode,
		Msg:  resp.Status,
	}, &docs
}

func (c *Client) GetTransactionByCustomerUUID(ledger string, collection string, uuid string) (*Response, *model.SearchResponse) {

	const path = "ics/api/v1/ledger"
	const contentType = "application/json"

	reqURL := fmt.Sprintf("%s/%s/%s/collection/%s/documents/search", c.URL, path, ledger, collection)

	qry := query.Query{
		Expressions: []query.Expression{
			{
				[]query.Constraint{
					{
						Field:    "uuid",
						Operator: query.EQUAL,
						Value:    uuid,
					},
				},
			},
		},
		OrderBy: []query.OrderConstraint{
			{
				Field: "_vault_md.ts",
				Order: query.DESC,
			},
		},
		Limit: 0,
	}

	search := query.Search{
		KeepOpen: true,
		Query:    qry,
		Page:     1,
		PerPage:  100,
	}

	data, err := json.Marshal(search)

	if err != nil {
		return &Response{Code: http.StatusInternalServerError, Msg: err.Error()}, nil
	}

	req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(data))

	if err != nil {
		return &Response{Code: http.StatusInternalServerError, Msg: err.Error()}, nil
	}

	req.Header.Set("accept", ": */*")
	req.Header.Set("X-API-Key", c.ApiKey)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return &Response{Code: http.StatusInternalServerError, Msg: err.Error()}, nil
	}
	defer resp.Body.Close()

	// read response body
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	var docs model.SearchResponse
	if err := json.Unmarshal(body, &docs); err != nil {
		return &Response{Code: http.StatusInternalServerError, Msg: err.Error()}, nil
	}

	return &Response{
		Code: resp.StatusCode,
		Msg:  resp.Status,
	}, &docs

}
