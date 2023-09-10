package mysql

import (
	"gangbu/pkg/models"
	"gorm.io/gorm"
)

type gameHistoryRepository struct {
}

func (ghr *gameHistoryRepository) CreateGame(gh *models.GameHistory, db *gorm.DB) error {
	return db.Create(gh).Error
}

func NewGameHistoryRepository() models.GameHistoryRepository {
	return &gameHistoryRepository{}
}
