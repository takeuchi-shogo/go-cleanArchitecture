package database

import (
	"errors"
	"time"

	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type UserTokenRepository struct{}

func (r *UserTokenRepository) FindByToken(db *gorm.DB, token string) (userToken domain.UserTokens, err error) {
	userToken = domain.UserTokens{}
	db.Where("token = ?", token).Last(&userToken)

	if userToken.ID == 0 {
		return domain.UserTokens{}, errors.New("user token not found")
	}

	return userToken, nil
}

func (r *UserTokenRepository) Create(db *gorm.DB, userToken domain.UserTokens) (newUserToken domain.UserTokens, err error) {
	newUserToken = domain.UserTokens{}

	var token string

	for {
		token = newUserToken.GetToken()
		u, err := r.FindByToken(db, token)

		if err != nil {
			break
		}

		if u.TokenExpireAt < time.Now().Unix() {
			break
		}
	}

	newUserToken.UserID = userToken.UserID
	newUserToken.Token = token
	newUserToken.RefreshToken = newUserToken.GetToken()
	newUserToken.CreatedAt = time.Now().Unix()
	newUserToken.SetExpireAt()

	db.NewRecord(newUserToken)
	err = db.Create(&newUserToken).Error

	return newUserToken, err
}
