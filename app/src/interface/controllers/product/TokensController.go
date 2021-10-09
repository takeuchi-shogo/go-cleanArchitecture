package product

import (
	"sns-sample/src/domain"
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/usecase/product"
)

type TokensController struct {
	Interactor product.UserTokenInteractor
}

func NewTokensController(db database.DB) *TokensController {
	return &TokensController{
		Interactor: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
	}
}

func (controller *TokensController) Post(c controllers.Context) {
	screenName := c.PostForm("id")
	password := c.PostForm("password")

	token, res := controller.Interactor.Create(domain.Users{
		ScreenName: screenName,
		Password:   password,
	})

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	c.JSON(201, controllers.NewH("success", token))
}

func (controller *TokensController) PostRefresh(c controllers.Context) {
	accessToken := c.PostForm("screenName")
	refreshToken := c.PostForm("refreshToken")

	token, res := controller.Interactor.Refresh(domain.UserTokens{
		Token:        accessToken,
		RefreshToken: refreshToken,
	})

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	c.JSON(201, controllers.NewH("success", token))
}
