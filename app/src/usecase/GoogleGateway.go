package usecase

type GoogleGateway interface {
	GetClientID() (clientID string)
	GetUserID(code string) (googleUserID string, err error)
}
