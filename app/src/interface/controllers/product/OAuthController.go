package product

import (
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/interface/gateways"
	"sns-sample/src/usecase/product"
)

type OAuthController struct {
	Interactor product.OAuthInteractor
}

type OAuthControllerProvider struct {
	DB     database.DB
	Google gateways.Google
}

func NewOAuthController(p OAuthControllerProvider) *OAuthController {
	return &OAuthController{
		Interactor: product.OAuthInteractor{
			DB:     &database.DBRepository{DB: p.DB},
			Google: &gateways.GoogleGateway{Google: p.Google},
		},
	}
}

func (controller *OAuthController) GetGoogle(c controllers.Context) {

	google, res := controller.Interactor.GetByGoogle()
	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, controllers.NewH("success", google))
}
