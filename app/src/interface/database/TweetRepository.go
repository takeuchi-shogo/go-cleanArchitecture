package database

import (
	"errors"
	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type TweetRepository struct{}

func (repo *TweetRepository) FindAll(db *gorm.DB) (tweets []domain.Tweets, err error) {
	tweets = []domain.Tweets{}
	db.Order("ID desc").Find(&tweets)
	if len(tweets) == 0 {
		return []domain.Tweets{}, errors.New("まだ呟いていません")
	}
	return tweets, nil
}

func (repo *TweetRepository) FindById(db *gorm.DB, id int) (tweet domain.Tweets, err error) {
	tweet = domain.Tweets{}
	db.First(&tweet, id)
	if tweet.ID <= 0 {
		return domain.Tweets{}, errors.New("この呟きは有りません")
	}
	return tweet, nil
}

func (repo *TweetRepository) Create(db *gorm.DB, tweet domain.Tweets) (newTweet domain.Tweets, err error) {
	newTweet = domain.Tweets{}
	newTweet.Contents = tweet.Contents

	db.NewRecord(newTweet)
	err = db.Create(&newTweet).Error

	return newTweet, err
}

func (repo *TweetRepository) Save(db *gorm.DB, tweet domain.TweetsForPatch) (t domain.Tweets, err error) {
	foundTweet, err := repo.FindById(db, tweet.ID)

	if err != nil {
		return domain.Tweets{}, err
	}

	foundTweet.Contents = tweet.Contents

	err = db.Save(&foundTweet).Error

	return foundTweet, err
}

func (repo *TweetRepository) Delete(db *gorm.DB, id int) (err error) {
	foundTweet, err := repo.FindById(db, id)

	if err != nil {
		return err
	}

	return db.Delete(&foundTweet).Error
}
