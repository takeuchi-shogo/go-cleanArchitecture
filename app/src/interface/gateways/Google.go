package gateways

type Google interface {
	GetClientID() (clientID string)
	GetUserID(code string) (userID string, err error)
}
