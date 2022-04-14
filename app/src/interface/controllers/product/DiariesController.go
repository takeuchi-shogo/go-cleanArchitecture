package product

import (
	"sns-sample/src/interface/controllers"
	"sns-sample/src/interface/database"
	"sns-sample/src/usecase/product"
	"strconv"
)

type DiariesController struct {
	Token      product.UserTokenInteractor
	Interactor product.DiaryInteractor
}

func NewDiariesController(db database.DB) *DiariesController {
	return &DiariesController{
		Token: product.UserTokenInteractor{
			DB:        &database.DBRepository{DB: db},
			UserToken: &database.UserTokenRepository{},
		},
		Interactor: product.DiaryInteractor{
			DB:    &database.DBRepository{DB: db},
			Diary: &database.DiaryRepository{},
		},
	}
}

func (controller *DiariesController) Get(c controllers.Context) {
	token, res := controller.Token.Verify(c.Query("accessToken"))

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	diaryID, _ := strconv.Atoi(c.Param("id"))

	diary, res := controller.Interactor.Get(token.UserID, diaryID)

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	c.JSON(res.StatusCode, controllers.NewH("success", diary))
}

func (controller *DiariesController) GetList(c controllers.Context) {
	token, res := controller.Token.Verify(c.Query("accessToken"))

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	diaries, res := controller.Interactor.GetList(token.UserID)

	if res.Error != nil {
		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
		return
	}

	c.JSON(res.StatusCode, controllers.NewH("success", diaries))
}

// func (controller *DiariesController) Post(c controllers.Context) {
// 	token, res := controller.Token.Verify(c.PostForm("accessToken"))
// 	if res.Error != nil {
// 		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
// 		return
// 	}

// 	title := c.PostForm("title")
// 	content := c.PostForm("content")

// 	diary := domain.Diaries{
// 		UserID:  token.UserID,
// 		Title:   title,
// 		Content: content,
// 	}

// 	newDiary, res := controller.Interactor.Create(diary)
// 	if res.Error != nil {
// 		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
// 		return
// 	}

// 	c.JSON(res.StatusCode, controllers.NewH("success", newDiary))
// }

// func (controller *DiariesController) Patch(c controllers.Context) {
// 	token, res := controller.Token.Verify(c.PostForm("accessToken"))
// 	if res.Error != nil {
// 		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
// 		return
// 	}

// 	title := c.PostForm("title")
// 	content := c.PostForm("content")

// 	diary := domain.Diaries{
// 		UserID:  token.UserID,
// 		Title:   title,
// 		Content: content,
// 	}

// 	newDiary, res := controller.Interactor.Patch(diary)
// 	if res.Error != nil {
// 		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
// 		return
// 	}

// 	c.JSON(res.StatusCode, controllers.NewH("success", newDiary))
// }

// func (controller *DiariesController) Delete(c controllers.Context) {
// 	token, res := controller.Token.Verify(c.PostForm("accessToken"))
// 	if res.Error != nil {
// 		c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
// 		return
// 	}

// 	diaryID := c.Param("id")

// 	if res := controller.Interactor.Delete(token.UserID, diaryID); res.Error != nil {
// 		if res.Error != nil {
// 			c.JSON(res.StatusCode, controllers.NewH(res.Error.Error(), nil))
// 			return
// 		}
// 	}

// 	c.JSON(res.StatusCode, controllers.NewH("success", nil))
// }
