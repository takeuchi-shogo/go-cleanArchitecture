package product

import (
	"sns-sample/src/domain"
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/usecase/product"
	"strconv"
)

type UsersController struct {
	Token      product.UserTokenInteractor
	Interactor product.UsersInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Token: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.UsersInteractor{
			DB:        &database.DBRepository{DB: db},
			User:      &database.UserRepository{},
			UserToken: &database.UserTokenRepository{},
		},
	}
}

func (controller *UsersController) Get(c controllers.Context) {
	token, res := controller.Token.Verify(c.Query("accessToken"))

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	user, res := controller.Interactor.Get(token.UserID)

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	c.JSON(res.StatusCode, controllers.NewH("success", user))
}

func (controller *UsersController) Post(c controllers.Context) {
	userName := c.PostForm("userName")
	sex := c.PostForm("sex")
	age, _ := strconv.Atoi(c.PostForm("age"))
	prefecture := c.PostForm("prefecture")
	city := c.PostForm("city")
	tel := c.PostForm("tel")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := domain.Users{
		UserName:   userName,
		Sex:        sex,
		Age:        age,
		Prefecture: &prefecture,
		City:       &city,
		Tel:        &tel,
		Email:      &email,
		Password:   password,
	}

	createUser, res := controller.Interactor.Create(user)

	if res.Error != nil {
		//　ログに書き込む際にパスワードは ***** にする
		user.Password = "*****"
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	c.JSON(res.StatusCode, controllers.NewH("success", createUser))
}

/*
func (controller *UsersController) Patch(c controllers.Context) {

}
*/
