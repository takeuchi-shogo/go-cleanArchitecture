package gateways

type GoogleGateway struct {
	Google Google
}

func (gateway *GoogleGateway) GetClientID() (clientID string) {
	return gateway.Google.GetClientID()
}

func (gateway *GoogleGateway) GetUserID(code string) (clientID string, err error) {
	return gateway.Google.GetUserID(code)
}
