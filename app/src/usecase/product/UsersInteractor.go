package product

import (
	"sns-sample/src/domain"

	"sns-sample/src/usecase"
)

type UsersInteractor struct {
	DB   usecase.DBRepository
	User usecase.UserRepository
}

func (interactor *UsersInteractor) Create(user domain.Users) (res domain.UserForMe, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	if _, err := interactor.User.FindByUserName(db, user.UserName); err != nil {
		return domain.UserForMe{}, usecase.NewResultStatus(400, err)
	}

	res, err := interactor.User.Create(db, user)

	if err != nil {
		return domain.UserForMe{}, usecase.NewResultStatus(400, err)
	}

	return res, usecase.NewResultStatus(201, nil)
}
