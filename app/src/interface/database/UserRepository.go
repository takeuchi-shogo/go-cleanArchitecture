package database

import (
	"errors"
	"sns-sample/src/domain"
	"time"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB DB
}

func (repo *UserRepository) FindById(db *gorm.DB, id int) (user domain.Users, err error) {
	user = domain.Users{}

	db.Where("id = ?", id).First(&user)

	if user.ID <= 0 {
		return domain.Users{}, errors.New("user is not found")
	}

	return user, nil
}

func (repo *UserRepository) FindByScreenName(db *gorm.DB, screenName string) (user domain.Users, err error) {
	user = domain.Users{}
	db.Where("screen_name = ?", screenName).First(&user)
	if user.ID == 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) FindByUserName(db *gorm.DB, userName string) (user domain.Users, err error) {
	user = domain.Users{}
	db.Where("user_name = ?", userName).First(&user)
	if user.ID == 0 {
		return domain.Users{}, errors.New("user is not found")
	}
	return user, nil
}

func (repo *UserRepository) Create(db *gorm.DB, user domain.Users) (newUser domain.Users, err error) {
	newUser = domain.Users{}
	newUser.ScreenName = user.ScreenName
	newUser.UserName = user.UserName
	newUser.Sex = user.Sex
	newUser.Age = user.Age
	newUser.Prefecture = user.Prefecture
	newUser.City = user.City
	newUser.Tel = user.Tel
	newUser.Email = user.Email
	newUser.Password = user.Password
	newUser.CreatedAt = time.Now().Unix()
	newUser.UpdatedAt = time.Now().Unix()

	newUser.Password = user.GetCodedPassword(newUser.Password)

	db.NewRecord(newUser)
	err = db.Create(&newUser).Error

	return newUser, err
}
