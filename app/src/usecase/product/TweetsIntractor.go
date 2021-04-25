package product

import (
	"errors"
	"sns-sample/src/domain"
	"sns-sample/src/usecase"
)

type TweetsInteractor struct {
	DB    usecase.DBRepository
	Tweet usecase.TweetRepository
}

type TweetList struct {
	Lists []domain.Tweets
}

func (interactor *TweetsInteractor) GetList() (tweetList TweetList, resultStatus *usecase.ResultStatus) {
	db := interactor.DB.Connect()

	tweets := []domain.Tweets{}

	tweets, err := interactor.Tweet.FindAll(db)

	if err != nil {
		return TweetList{Lists: tweets}, usecase.NewResultStatus(404, err)
	}

	/*
		for _, list := range lists {
			tweet, err := interactor.build(db, list)
			if err != nil {
				continue
			}
			tweets = append(tweets, tweet)
		}
	*/

	return TweetList{Lists: tweets}, usecase.NewResultStatus(200, nil)
}

func (interactor *TweetsInteractor) Get(id int) (tweet domain.Tweets, resultStatus *usecase.ResultStatus) {
	db := interactor.DB.Connect()

	foundTweet, err := interactor.Tweet.FindById(db, id)

	if err != nil {
		return domain.Tweets{}, usecase.NewResultStatus(404, errors.New(""))
	}
	return foundTweet, usecase.NewResultStatus(200, nil)
}

func (interactor *TweetsInteractor) Create(tweet domain.Tweets) (createTweet domain.Tweets, resultStatus *usecase.ResultStatus) {
	db := interactor.DB.Connect()

	createTweet, err := interactor.Tweet.Create(db, tweet)

	if err != nil {
		return domain.Tweets{}, usecase.NewResultStatus(404, errors.New(""))
	}
	return createTweet, usecase.NewResultStatus(200, nil)
}

func (interactor *TweetsInteractor) Save(newTweet domain.TweetsForPatch) (tweet domain.Tweets, resultStatus *usecase.ResultStatus) {
	db := interactor.DB.Connect()

	patchTweet, err := interactor.Tweet.Save(db, newTweet)

	if err != nil {
		return domain.Tweets{}, usecase.NewResultStatus(404, errors.New(""))
	}
	return patchTweet, usecase.NewResultStatus(200, nil)
}

func (interactor *TweetsInteractor) Delete(id int) (resultStatus *usecase.ResultStatus) {
	db := interactor.DB.Connect()

	err := interactor.Tweet.Delete(db, id)

	if err != nil {
		return usecase.NewResultStatus(404, errors.New(""))
	}

	return usecase.NewResultStatus(200, nil)
}

/*
func (interactor *TweetsInteractor) build(db *gorm.DB, list domain.Tweets) (tweet domain.Tweets, err error) {
	tweets,
}
*/
