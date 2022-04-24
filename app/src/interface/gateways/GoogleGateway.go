package gateways

type GoogleGateway struct {
	Google Google
}

func (gateway *GoogleGateway) GetLoginURL(state string) (clientID string) {
	return gateway.Google.GetLoginURL(state)
}

func (gateway *GoogleGateway) GetUserID(code string) (clientID string, err error) {
	return gateway.Google.GetUserID(code)
}
