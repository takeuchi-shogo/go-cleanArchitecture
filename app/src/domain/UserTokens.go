package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type UserTokens struct {
	ID                   int    `json:"id"`
	UserID               int    `json:"adminUserId"`
	Token                string `json:"toekn"`
	TokenExpireAt        int64  `json:"toeknExpireAt"`
	RefreshToken         string `json:"refreshToken"`
	RefreshTokenExpireAt int64  `json:"refreshTokenExpireAt"`
	CreatedAt            int64  `json:"createAt"`
}

type UserTokensForGet struct {
	Token                string `json:"token"`
	TokenExpireAt        int64  `json:"tokenExpireAt"`
	RefreshToken         string `json:"refreshToken"`
	RefreshTokenExpireAt int64  `json:"refreshTokenExpireAt"`
}

/* func (u *UserTokens) Validate() (err error) {
	if err := u.checkUserID(); err != nil { return err }
	if err := u.checkToken(); err != nil { return err }
	if err := u.checkTokenExpireAt(); err != nil { return err }
	if err := u.checkRefreshToken(); err != nil { return err }
	if err := u.checkRefreshTokenExpireAt(); err != nil { return err }
	return nil
} */

func (u *UserTokens) GetToken() string {

	source := rand.NewSource(time.Now().UnixNano())

	maxRange := 50
	minRange := 40

	str := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	cnt := len(str)

	var token string

	for i := 0; i < maxRange; i++ {
		if minRange <= i {
			if rand.New(source).Intn(maxRange-minRange) == 0 {
				break
			}
		}
		token = token + fmt.Sprintf("%c", str[rand.New(source).Intn(cnt)])
	}

	return token
}

func (u *UserTokens) SetExpireAt() {
	u.TokenExpireAt = time.Now().Unix() + (60 * 60 * 24 * 3)
	u.RefreshTokenExpireAt = time.Now().Unix() + (60 * 60 * 24 * 30)
}

func (u *UserTokens) BuildForGet() UserTokensForGet {
	token := UserTokensForGet{}
	token.Token = u.Token
	token.TokenExpireAt = u.TokenExpireAt
	token.RefreshToken = u.RefreshToken
	token.RefreshTokenExpireAt = u.RefreshTokenExpireAt
	return token
}
