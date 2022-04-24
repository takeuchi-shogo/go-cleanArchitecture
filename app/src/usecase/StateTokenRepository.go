package usecase

import (
	"sns-sample/src/domain"

	"github.com/jinzhu/gorm"
)

type StateTokenRepository interface {
	Create(db *gorm.DB, stateToken domain.StateTokens) (newStateToken domain.StateTokens, err error)
}
