package domain

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
)

// DB 保存する構造体
type Users struct {
	ID         int    `json:"id"`
	ScreenName string `json:"screenName"`
	UserName   string `json:"userName"`
	//LastName   string  `json:"lastName"`
	//FirstName  string  `json:"firstName"`
	Sex        string  `json:"sex"`
	Age        int     `json:"age"`
	Prefecture *string `json:"addressPrefecture"`
	City       *string `json:"addressCity"`
	Tel        *string `json:"tel"`
	//IsAuthorizedTel   bool    `json:"isAuthorizedTel"`
	Email *string `json:"email"`
	//IsAuthorizedEmail bool    `json:"isAuthorizedEmail"`
	Password string `json:"password"`
	//LastLoginAt       int64   `json:"lastLoginAt"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

type UserForMe struct {
	ID         int     `json:"id"`
	UserName   string  `json:"userName"`
	Sex        string  `json:"sex"`
	Age        int     `json:"age"`
	Prefecture *string `json:"addressPrefecture"`
	City       *string `json:"addressCity"`
	Tel        *string `json:"tel"`
	Email      *string `json:"email"`
	Password   string  `json:"password"`
}

//編集時のユーザーの構造体
type UserForPatch struct {
	ID                int     `json:"id"`
	ScreenName        string  `json:"screenName"`
	UserName          string  `json:"userName"`
	Sex               string  `json:"sex"`
	Age               int     `json:"age"`
	Prefecture        *string `json:"addressPrefecture"`
	City              *string `json:"addressCity"`
	Tel               *string `json:"tel"`
	IsAuthorizedTel   bool
	Email             *string `json:"email"`
	IsAuthorizedEmail *bool
	Password          string `json:"password"`
}

//初期値としてランダムにログインIDを生成する
func (u *Users) RandScreenName(digit uint32) (string, error) {
	const letters = "abcdefghijklnmopqrstuvwxyzABCDEFGHIJKLNMOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}

	return result, nil
}

func (u *Users) GetCodedPassword(plainPassword string) string {
	var data [sha256.Size]byte
	data = sha256.Sum256([]byte(plainPassword))
	return fmt.Sprintf("%x", data)
}
