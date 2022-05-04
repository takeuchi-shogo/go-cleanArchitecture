package usecase

import (
	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type UserTokenRepository interface {
	FindByToken(db *gorm.DB, token string) (userToken domain.UserTokens, err error)
	Create(db *gorm.DB, userToken domain.UserTokens) (newUserToken domain.UserTokens, err error)
}
