package product

import (
	"sns-sample/src/domain"
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/interface/gateways"
	"sns-sample/src/usecase/product"
)

type TokensController struct {
	Interactor product.UserTokenInteractor
}

type TokensControllerProvider struct {
	DB     database.DB
	Google gateways.Google
}

func NewTokensController(p TokensControllerProvider) *TokensController {
	return &TokensController{
		Interactor: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: p.DB},
			Google:    &gateways.GoogleGateway{Google: p.Google},
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

func (controller *TokensController) PostGoogle(c controllers.Context) {
	state := c.PostForm("state")
	code := c.PostForm("code")
	token, res := controller.Interactor.CreateByGoogle(state, code)
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
