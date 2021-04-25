package product

import (
	"sns-sample/src/domain"
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/usecase/product"
	"strconv"
)

type UsersController struct {
	Interactor product.UsersInteractor
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: product.UsersInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

/*
func (controller *UsersController) Get(c controllers.Context) {

}
*/

func (controller *UsersController) Post(c controllers.Context) {
	userName := c.PostForm("UserName")
	sex := c.PostForm("Sex")
	age, _ := strconv.Atoi(c.PostForm("Age"))
	addressPrefecture := c.PostForm("AddressPrefecture")
	addressCity := c.PostForm("AddressCity")
	tel := c.PostForm("Tel")
	email := c.PostForm("Email")
	password := c.PostForm("Password")

	user := domain.Users{
		UserName:          userName,
		Sex:               sex,
		Age:               age,
		AddressPrefecture: &addressPrefecture,
		AddressCity:       &addressCity,
		Tel:               &tel,
		Email:             &email,
		Password:          password,
	}

	createUser, res := controller.Interactor.Create(user)

	if res.Error != nil {
		//　ログに書き込む際にパスワードは ***** にする
		user.Password = "*****"
	}

	c.JSON(res.StatusCode, controllers.NewH("success", createUser))
}

/*
func (controller *UsersController) Patch(c controllers.Context) {

}
*/
