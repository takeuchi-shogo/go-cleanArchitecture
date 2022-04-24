package database

import (
	"sns-sample/src/domain"
	"time"

	"github.com/jinzhu/gorm"
)

type StateTokenRepository struct{}

func (repo *StateTokenRepository) Create(db *gorm.DB, stateToken domain.StateTokens) (newStateToken domain.StateTokens, err error) {
	newStateToken = domain.StateTokens{}

	newStateToken.ApplicationName = stateToken.ApplicationName
	newStateToken.StateToken = stateToken.GetToken()
	newStateToken.CreatedAt = time.Now().Unix()
	newStateToken.SetExpireAt()

	if err = newStateToken.Validate(); err != nil {
		return domain.StateTokens{}, err
	}

	return newStateToken, nil
}
