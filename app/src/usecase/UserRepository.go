package usecase

import (
	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	FindById(db *gorm.DB, id int) (user domain.Users, err error)
	FindByScreenName(db *gorm.DB, screenName string) (user domain.Users, err error)
	FindByUserName(db *gorm.DB, Name string) (user domain.Users, err error)
	FindByPrefecture(db *gorm.DB, prefecture string) (users []domain.Users, err error)
	Create(db *gorm.DB, user domain.Users) (res domain.Users, err error)
}
