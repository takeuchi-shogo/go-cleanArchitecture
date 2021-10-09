package product

import (
	"sns-sample/src/domain"

	"sns-sample/src/usecase"
)

type UsersInteractor struct {
	DB        usecase.DBRepository
	User      usecase.UserRepository
	UserToken usecase.UserTokenRepository
}

func (interactor *UsersInteractor) Get(id int) (res domain.Users, resultStatus *usecase.ResultStatus) {
	db := interactor.DB.Connect()

	res, err := interactor.User.FindById(db, id)

	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, err)
	}

	return res, usecase.NewResultStatus(200, nil)
}

func (interactor *UsersInteractor) Create(user domain.Users) (res domain.Users, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	/* if _, err := interactor.User.FindByUserName(db, user.UserName); err != nil {
		return domain.UserForMe{}, usecase.NewResultStatus(400, err)
	} */

	randScreenName, err := user.RandScreenName(12)

	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, err)
	}

	user.ScreenName = randScreenName

	res, err = interactor.User.Create(db, user)

	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, err)
	}

	newUserToken := domain.UserTokens{}
	newUserToken.UserID = res.ID

	_, err = interactor.UserToken.Create(db, newUserToken)

	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, err)
	}

	return res, usecase.NewResultStatus(201, nil)
}
