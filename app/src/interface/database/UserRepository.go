package database

import (
	"errors"
	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB DB
}

func (repo *UserRepository) FindByUserName(db *gorm.DB, userName string) (user domain.Users, err error) {
	user = domain.Users{}
	db.Where("user_name = ?", userName).First(&user)
	if user.ID == 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, user domain.Users) (newUser domain.UserForMe, err error) {
	newUser = domain.UserForMe{}
	newUser.UserName = user.UserName
	newUser.Sex = user.Sex
	newUser.Age = user.Age
	newUser.AddressPrefecture = user.AddressPrefecture
	newUser.AddressCity = user.AddressCity
	newUser.Tel = user.Tel
	newUser.Email = user.Email
	newUser.Password = user.Password

	db.NewRecord(newUser)
	err = db.Create(&newUser).Error

	return newUser, err
}
