package product

import (
	"sns-sample/src/domain"
	"sns-sample/src/usecase"
)

type DiaryInteractor struct {
	DB    usecase.DBRepository
	Diary usecase.DiaryRepository
}

func (interactor *DiaryInteractor) Get(userID, diaryID int) (diary domain.Diaries, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	diary, err := interactor.Diary.FindByIDAndUserID(db, diaryID, userID)

	if err != nil {
		return domain.Diaries{}, usecase.NewResultStatus(404, err)
	}

	return diary, usecase.NewResultStatus(200, nil)
}

func (interactor *DiaryInteractor) GetList(userID int) (diaries []domain.Diaries, resultStatus *usecase.ResultStatus) {

	db := interactor.DB.Connect()

	diaries, err := interactor.Diary.FindByUserID(db, userID)

	if err != nil {
		return []domain.Diaries{}, usecase.NewResultStatus(404, err)
	}

	return diaries, usecase.NewResultStatus(200, nil)
}

// func (interactor *DiaryInteractor) Create(diary domain.Diaries) (newDiary domain.Diaries, resultStatus *usecase.ResultStatus) {

// 	db := interactor.DB.Connect()

// 	newDiary, err := interactor.Diary.Create(db, diary)
// 	if err != nil {
// 		return domain.Diaries{}, usecase.NewResultStatus(400, err)
// 	}
// 	return newDiary, usecase.NewResultStatus(200, nil)
// }

// func (interactor *DiaryInteractor) Patch(diary domain.Diaries) (newDiary domain.Diaries, resultStatus *usecase.ResultStatus) {

// }

// func (interactor *DiaryInteractor) Delete(userID, diaryID int) (resultStatus *usecase.ResultStatus) {
// 	db := interactor.DB.Connect()

// 	err := interactor.Diary.Delete(db, diaryID)
// 	if err != nil {
// 		return usecase.NewResultStatus(400, err)
// 	}
// 	return usecase.NewResultStatus(200, nil)
// }
