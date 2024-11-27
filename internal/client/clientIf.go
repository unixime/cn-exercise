package client

type ImmuDBClientIf interface {
	Login(host string, port int, user string, password string) error
	Logout() error
	IsConnected() bool
	CreateDatabase(name string) error
	CreateTable(tableName string, tableDef string) error
}
