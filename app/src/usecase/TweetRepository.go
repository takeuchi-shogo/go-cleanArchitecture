package usecase

import (
	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type TweetRepository interface {
	FindAll(db *gorm.DB) (tweets []domain.Tweets, err error)
	FindById(db *gorm.DB, id int) (tweet domain.Tweets, err error)
	Create(db *gorm.DB, createTweet domain.Tweets) (tweet domain.Tweets, err error)
	Save(db *gorm.DB, newTweet domain.TweetsForPatch) (tweet domain.Tweets, err error)
	Delete(db *gorm.DB, id int) (err error)
}
