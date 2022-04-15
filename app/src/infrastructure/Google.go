package infrastructure

import (
	"context"
	"errors"
	"fmt"

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
			Scopes:       []string{"openid"},
			RedirectURL:  "http://localhost:8080/auth/callback/google",
		},
	}

	if google.Config == nil {
		panic("==== invalid key. google api ====")
	}

	return google
}

func (g *Google) GetClientID() (clientID string) {
	return g.Config.ClientID
}

func (g *Google) GetUserID(code string) (googleUserID string, err error) {

	fmt.Println("aaaaa: ", g.Config.AuthCodeURL(""))

	cxt := context.Background()

	httpClient, _ := g.Config.Exchange(cxt, code)
	if httpClient == nil {
		return "", errors.New("接続エラー")
	}

	client := g.Config.Client(cxt, httpClient)
	// userInfo, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	// if err != nil {
	// 	return "", errors.New("接続エラー")
	// }
	service, _ := v2.New(client)
	if err != nil {
		return "", errors.New("接続エラー")
	}

	userInfo, _ := service.Tokeninfo().AccessToken(httpClient.AccessToken).Context(cxt).Do()
	if err != nil {
		return "", errors.New("接続エラー")
	}

	// user, err := service.TokenInfo().AccessToken(httpClient.AccessToken).Context().Do()

	return userInfo.UserId, nil
}
