package product

import (
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/usecase/product"
)

type MeController struct {
	Token      product.UserTokenInteractor
	Interactor product.MeInteractor
}

func NewMeController(db database.DB) *MeController {
	return &MeController{
		Token: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.MeInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

func (controller *MeController) Get(c controllers.Context) {
	token, res := controller.Token.Verify(c.Query("accessToken"))
	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	me, res := controller.Interactor.Get(token.UserID)
	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, controllers.NewH("success", me))

}
