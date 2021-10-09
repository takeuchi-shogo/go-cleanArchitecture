package product

import (
	"sns-sample/src/domain"
	"sns-sample/src/usecase"
)

type UserSearchInteracor struct {
	DB   usecase.DBRepository
	User usecase.UserRepository
}

func (interactor *UserSearchInteracor) Get(userId int) (user domain.Users, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	user, err := interactor.User.FindById(db, userId)

	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, err)
	}

	return user, usecase.NewResultStatus(200, nil)
}

func (interactor *UserSearchInteracor) GetList(userId int) (users []domain.Users, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	user, err := interactor.User.FindById(db, userId)

	if err != nil {
		return []domain.Users{}, usecase.NewResultStatus(400, err)
	}

	users, err = interactor.User.FindByPrefecture(db, *user.Prefecture)

	if err != nil {
		return []domain.Users{}, usecase.NewResultStatus(400, err)
	}
	return users, usecase.NewResultStatus(200, nil)
}
