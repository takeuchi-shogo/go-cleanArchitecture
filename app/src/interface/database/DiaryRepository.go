package database

import (
	"errors"
	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type DiaryRepository struct{}

func (d *DiaryRepository) FindByID(db *gorm.DB, id int) (foundDiary domain.Diaries, err error) {
	foundDiary = domain.Diaries{}
	db.Where("id = ?", id).First(&foundDiary)
	if foundDiary.ID == 0 {
		return domain.Diaries{}, errors.New("日記が見つかりません")
	}
	return foundDiary, nil
}

func (d *DiaryRepository) FindByUserID(db *gorm.DB, userID int) (foundDiaries []domain.Diaries, err error) {
	foundDiaries = []domain.Diaries{}
	db.Where("user_id = ?", userID).Find(&foundDiaries)
	if len(foundDiaries) == 0 {
		return []domain.Diaries{}, errors.New("まだ日記を投稿していません")
	}
	return foundDiaries, nil
}

func (d *DiaryRepository) FindByIDAndUserID(db *gorm.DB, id, userID int) (foundDiary domain.Diaries, err error) {
	foundDiary = domain.Diaries{}
	db.Where("id = ? and user_id = ?", id, userID).Find(&foundDiary)
	if foundDiary.ID == 0 {
		return domain.Diaries{}, errors.New("日記が見つかりません")
	}
	return foundDiary, nil
}

// func (d *DiaryRepository) Create(db *gorm.DB, diary domain.Diaries) (newDiary domain.Diaries, err error) {

// }

// func (d *DiaryRepository) Delete(db *gorm.DB, id int) (err error) {

// }
