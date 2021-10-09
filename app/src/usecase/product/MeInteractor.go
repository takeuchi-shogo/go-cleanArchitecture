package product

import (
	"sns-sample/src/domain"
	"sns-sample/src/usecase"
)

type MeInteractor struct {
	DB   usecase.DBRepository
	User usecase.UserRepository
}

func (interactor *MeInteractor) Get(userID int) (me domain.Users, resultStatus *usecase.ResultStatus) {
	db := interactor.DB.Connect()

	me, err := interactor.User.FindById(db, userID)

	if err != nil {
		return domain.Users{}, usecase.NewResultStatus(400, err)
	}

	return me, usecase.NewResultStatus(200, nil)
}
