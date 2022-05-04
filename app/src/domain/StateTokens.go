package domain

import (
	"errors"
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

func (s *StateTokens) Validate() error {
	if err := s.checkApplicationName(); err != nil {
		return err
	}
	if err := s.checkStateToken(); err != nil {
		return err
	}
	if err := s.checkExpireAt(); err != nil {
		return err
	}
	return nil
}

func (s *StateTokens) checkApplicationName() error {
	if s.ApplicationName == "" {
		return errors.New("Application Nameが空です")
	}
	return nil
}

func (s *StateTokens) checkStateToken() error {
	if s.StateToken == "" {
		return errors.New("State Token not valid")
	}
	return nil
}

func (s *StateTokens) checkExpireAt() error {
	if s.ExpireAt <= 0 {
		return errors.New("ExpireAt not valid")
	}
	return nil
}

func (s *StateTokens) SetExpireAt() {
	s.ExpireAt = time.Now().Unix() + (60 * 60)
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

func (s *StateTokens) BuildForGet() (stateToken StateTokensForGet) {
	stateToken = StateTokensForGet{}
	return stateToken
}
