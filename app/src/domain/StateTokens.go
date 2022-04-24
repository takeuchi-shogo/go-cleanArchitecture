package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type StateTokens struct {
	ID              int    `json:"id"`
	ApplicationName string `json:"applicationName"`
	StateToken      string `json:"stateToken"`
	ExpireAt        int64  `json:"expireAt"`
	CreatedAt       int64  `json:"createdAt"`
}

type StateTokensForGet struct {
	LoginURL string `json:"loginUrl"`
}

func (s *StateTokens) GetToken() string {

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
