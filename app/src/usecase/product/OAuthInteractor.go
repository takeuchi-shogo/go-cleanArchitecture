package product

import "sns-sample/src/usecase"

type OAuthInteractor struct {
	DB     usecase.DBRepository
	Google usecase.GoogleGateway
}

func (interactor *OAuthInteractor) GetByGoogle() (google string, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	googleStateToken, err := interactor.GoogleStateToken.Create(db)
	if err != nil {
		return "", usecase.NewResultStatus(400, err)
	}
	google = interactor.Google.GetLoginURL(googleStateToken.StateToken)

	return "", usecase.NewResultStatus(200, nil)
}
