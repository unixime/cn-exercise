package client

type ImmuDBClientIf interface {
	Login() error
	Logout() error
}
