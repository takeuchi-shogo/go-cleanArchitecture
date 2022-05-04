package product

import (
	"strconv"

	"sns-sample/src/domain"
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/usecase/product"
)

type TweetsController struct {
	Token      product.UserTokenInteractor
	Interactor product.TweetInteractor
}

func NewTweetsController(db database.DB) *TweetsController {
	return &TweetsController{
		Token: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.TweetInteractor{
			DB:    &database.DBRepository{DB: db},
			Tweet: &database.TweetRepository{},
		},
	}
}

func (controller *TweetsController) GetList(c controllers.Context) {
	tweets, res := controller.Interactor.GetList()
	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), tweets))
		return
	}
	c.JSON(res.StatusCode, controllers.NewH("success", tweets))
}

func (controller *TweetsController) Get(c controllers.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tweet, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), tweet))
		return
	}
	c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), tweet))
}

func (controller *TweetsController) Post(c controllers.Context) {

	token, res := controller.Token.Verify(c.PostForm("accessToken"))

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	contents := c.PostForm("contents")

	createTweet, res := controller.Interactor.Create(domain.Tweets{
		UserID:   token.UserID,
		Contents: contents,
	})
	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, controllers.NewH("success", createTweet))
}

func (controller *TweetsController) Patch(c controllers.Context) {

	token, res := controller.Token.Verify(c.PostForm("accessToken"))

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	contents := c.PostForm("contents")

	updateTweet := domain.TweetsForPatch{
		ID:       id,
		UserID:   token.UserID,
		Contents: contents,
	}

	tweet, res := controller.Interactor.Save(updateTweet)
	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), tweet))
		return
	}
	c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), tweet))
}

func (controller *TweetsController) Delete(c controllers.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res := controller.Interactor.Delete(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
}
