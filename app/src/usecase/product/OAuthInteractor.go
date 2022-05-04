package product

import (
	"sns-sample/src/domain"
	"sns-sample/src/usecase"
)

type OAuthInteractor struct {
	DB         usecase.DBRepository
	Google     usecase.GoogleGateway
	StateToken usecase.StateTokenRepository
}

func (interactor *OAuthInteractor) GetByGoogle(stateToken domain.StateTokens) (buildStateToken domain.StateTokensForGet, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	googleStateToken, err := interactor.StateToken.Create(db, stateToken)
	if err != nil {
		return domain.StateTokensForGet{}, usecase.NewResultStatus(400, err)
	}
	loginURL := interactor.Google.GetLoginURL(googleStateToken.StateToken)

	buildStateToken = googleStateToken.BuildForGet()
	buildStateToken.LoginURL = loginURL
	return buildStateToken, usecase.NewResultStatus(200, nil)
}
