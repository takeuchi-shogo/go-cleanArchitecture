package usecase

import (
	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	FindByUserName(db *gorm.DB, Name string) (user domain.Users, err error)
	Create(db *gorm.DB, user domain.Users) (res domain.UserForMe, err error)
}
