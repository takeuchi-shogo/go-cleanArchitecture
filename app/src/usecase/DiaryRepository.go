package usecase

import (
	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type DiaryRepository interface {
	FindByID(db *gorm.DB, id int) (diary domain.Diaries, err error)
	FindByUserID(db *gorm.DB, userID int) (diaries []domain.Diaries, err error)
	FindByIDAndUserID(db *gorm.DB, id, userID int) (diaries domain.Diaries, err error)
	// Create(db *gorm.DB, diary domain.Diaries) (newDiary domain.Diaries, err error)
	// Delete(db *gorm.DB, id int) (err error)
}
