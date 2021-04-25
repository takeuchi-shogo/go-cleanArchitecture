package product

import (
	"sns-sample/src/domain"
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/usecase/product"
	"strconv"
)

type TweetsController struct {
	Interactor product.TweetsInteractor
}

func NewTweetsController(db database.DB) *TweetsController {
	return &TweetsController{
		Interactor: product.TweetsInteractor{
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
	contents := c.PostForm("Contents")

	tweet, res := controller.Interactor.Create(domain.Tweets{
		Contents: contents,
	})
	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), tweet))
		return
	}
	c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), tweet))
}

func (controller *TweetsController) Patch(c controllers.Context) {
	id, _ := strconv.Atoi(c.Param("tweetID"))
	contents := c.PostForm("contents")

	updateTweet := domain.TweetsForPatch{
		ID:       id,
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
