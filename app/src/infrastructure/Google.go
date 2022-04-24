package infrastructure

import (
	"context"
	"errors"

	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
	v2 "google.golang.org/api/oauth2/v2"
)

type Google struct {
	Config *oauth2.Config
}

func NewGoogle(c *Config) *Google {
	return newGoogle(c)
}

func newGoogle(c *Config) *Google {

	google := &Google{
		Config: &oauth2.Config{
			ClientID:     c.Google.ClientID,
			ClientSecret: c.Google.ClientSecret,
			Endpoint:     googleOAuth.Endpoint,
			Scopes: []string{
				"openid",
			},
			RedirectURL: "http://localhost:8080/auth/callback/google",
		},
	}

	if google.Config == nil {
		panic("==== invalid key. google api ====")
	}

	return google
}

func (g *Google) GetLoginURL(state string) (clientID string) {
	return g.Config.AuthCodeURL(state)
}

func (g *Google) GetUserID(code string) (googleUserID string, err error) {

	cxt := context.Background()

	httpClient, _ := g.Config.Exchange(cxt, code)
	if httpClient == nil {
		return "", errors.New("接続エラー")
	}

	client := g.Config.Client(cxt, httpClient)

	service, err := v2.New(client)
	if err != nil {
		return "", errors.New("接続エラー")
	}

	userInfo, err := service.Tokeninfo().AccessToken(httpClient.AccessToken).Context(cxt).Do()
	if err != nil {
		return "", errors.New("接続エラー")
	}

	return userInfo.UserId, nil
}
