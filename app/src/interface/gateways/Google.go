package gateways

type Google interface {
	GetLoginURL(state string) (clientID string)
	GetUserID(code string) (userID string, err error)
}
