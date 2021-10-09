package product

import (
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/usecase/product"
	"strconv"
)

type UserSearchesController struct {
	Token      product.UserTokenInteractor
	Interactor product.UserSearchInteracor
}

func NewUserSearchesController(db database.DB) *UserSearchesController {
	return &UserSearchesController{
		Token: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.UserSearchInteracor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

func (controller *UserSearchesController) Get(c controllers.Context) {
	_, res := controller.Token.Verify(c.Query("accessToken"))

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	userId, _ := strconv.Atoi(c.Param("userId"))
	user, res := controller.Interactor.Get(userId)

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	c.JSON(res.StatusCode, controllers.NewH("success", user))
}

func (controller *UserSearchesController) GetList(c controllers.Context) {
	token, res := controller.Token.Verify(c.Query("accessToken"))

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	users, res := controller.Interactor.GetList(token.UserID)

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	c.JSON(res.StatusCode, controllers.NewH("success", users))
}
