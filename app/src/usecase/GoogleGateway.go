package usecase

type GoogleGateway interface {
	GetLoginURL(state string) (clientID string)
	GetUserID(code string) (googleUserID string, err error)
}
