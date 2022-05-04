package product

import (
	"errors"
	"fmt"
	"sns-sample/src/domain"
	"sns-sample/src/usecase"
	"time"
)

type UserTokenInteractor struct {
	DB        usecase.DBRepository
	Google    usecase.GoogleGateway
	User      usecase.UserRepository
	UserToken usecase.UserTokenRepository
}

func (interactor *UserTokenInteractor) Verify(token string) (userToken domain.UserTokens, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	userToken, err := interactor.UserToken.FindByToken(db, token)

	if err != nil {
		return domain.UserTokens{}, usecase.NewResultStatus(401, errors.New("unauthorized"))
	}

	if userToken.TokenExpireAt < time.Now().Unix() {
		if time.Now().Unix() < userToken.RefreshTokenExpireAt {
			return domain.UserTokens{}, usecase.NewResultStatus(406, errors.New("need refresh a token"))
		}
		return domain.UserTokens{}, usecase.NewResultStatus(401, errors.New("unauthorized"))
	}

	return userToken, usecase.NewResultStatus(200, nil)
}

func (interactor *UserTokenInteractor) Create(user domain.Users) (createUserToken domain.UserTokensForGet, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	foundUser, err := interactor.User.FindByScreenName(db, user.ScreenName)

	if err != nil {
		return domain.UserTokensForGet{}, usecase.NewResultStatus(400, errors.New("Incorrect login ID or password"))
	}

	if foundUser.GetCodedPassword(user.Password) != foundUser.Password {
		return domain.UserTokensForGet{}, usecase.NewResultStatus(400, errors.New("Incorrect password"))
	}

	newUserToken := domain.UserTokens{}
	newUserToken.UserID = foundUser.ID

	res, err := interactor.UserToken.Create(db, newUserToken)

	if err != nil {
		return domain.UserTokensForGet{}, usecase.NewResultStatus(400, err)
	}

	return res.BuildForGet(), usecase.NewResultStatus(201, nil)
}

func (interactor *UserTokenInteractor) CreateByGoogle(state string, code string) (createdUserToken domain.UserTokensForGet, resultStatus *usecase.ResultStatus) {

	// db := interactor.DB.Connect()

	// _, err := interactor.GoogleToken.FindByState(db, state)
	// if err != nil {
	// 	return domain.UserTokensForGet{}, usecase.NewResultStatus(400, errors.New("認証されたユーザーではありません"))
	// }

	googleUserID, err := interactor.Google.GetUserID(code)
	if err != nil {
		return domain.UserTokensForGet{}, usecase.NewResultStatus(400, errors.New("Googleアカウントの認証が失敗しました"))
	}

	fmt.Println(googleUserID)

	// applicationName := domain.NewUserOauthCertificationsApplicationName()

	// userOauthCertification, err := interactor.UserOauthCertification.FindByApplication(db, applicationName.Google, googleUserID)
	// if err != nil {
	// 	return domain.UserTokensForGet{}, errors.New("Googleアカウントの認証が失敗しました")
	// }

	// newUserToken := domain.UserTokens{}
	// newUserToken.UserID = userOauthCertification.UserID

	// res, err := interactor.UserToken.Create(db, newUserToken)
	// if err != nil {
	// 	return domain.UserTokensForGet{}, usecase.NewResultStatus(400, err)
	// }
	// return res.BuildForGet(), usecase.NewResultStatus(201, nil)
	return domain.UserTokensForGet{}, usecase.NewResultStatus(201, nil)
}

func (interactor *UserTokenInteractor) Refresh(userToken domain.UserTokens) (createUserToken domain.UserTokensForGet, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	foundUserToken, err := interactor.UserToken.FindByToken(db, userToken.Token)

	if err != nil {
		return domain.UserTokensForGet{}, usecase.NewResultStatus(400, errors.New("could not refresh a token"))
	}
	if foundUserToken.RefreshToken != userToken.RefreshToken {
		return domain.UserTokensForGet{}, usecase.NewResultStatus(400, errors.New("could not refresh a token"))
	}
	if foundUserToken.RefreshTokenExpireAt < time.Now().Unix() {
		return domain.UserTokensForGet{}, usecase.NewResultStatus(400, errors.New("could not refresh a token"))
	}

	newUserToken := domain.UserTokens{}
	newUserToken.UserID = foundUserToken.UserID

	res, err := interactor.UserToken.Create(db, newUserToken)

	if err != nil {
		return domain.UserTokensForGet{}, usecase.NewResultStatus(400, err)
	}

	return res.BuildForGet(), usecase.NewResultStatus(200, nil)
}
