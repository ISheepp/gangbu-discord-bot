package mysql

import (
	"gangbu/pkg/e"
	"gangbu/pkg/models"
	"gorm.io/gorm"
)

type gameHistoryRepository struct {
}

func (ghr *gameHistoryRepository) GetLastFiveGameHistoryByDiscordId(discordId string, db *gorm.DB) ([]models.GameHistory, error) {
	var gh []models.GameHistory
	result := db.Where("player_discord_user_id = ? and game_status = ?", discordId, e.FINISHED).Order("created_at desc").Limit(5).Find(&gh)
	if result.Error != nil {
		return nil, result.Error
	}
	return gh, nil
}

func (ghr *gameHistoryRepository) GetGameHistoryByDiscordId(discordId string, db *gorm.DB) ([]models.GameHistory, error) {
	var gh []models.GameHistory
	result := db.Where("player_discord_user_id = ?", discordId).Find(&gh)
	if result.Error != nil {
		return nil, result.Error
	}
	return gh, nil
}

func (ghr *gameHistoryRepository) UpdateGameAfterMainBet(game *models.GameHistory, db *gorm.DB) error {
	return db.Save(game).Error
}

func (ghr *gameHistoryRepository) GetGameHistoryByRequestId(requestId string, db *gorm.DB) (*models.GameHistory, error) {
	var gh models.GameHistory
	result := db.Where("request_id = ?", requestId).First(&gh)
	if result.Error != nil {
		return nil, result.Error
	}
	return &gh, nil
}

func (ghr *gameHistoryRepository) UpdateRequestIdByTxId(txId string, requestID string, db *gorm.DB) error {
	result := db.Model(&models.GameHistory{}).Where("request_random_tx_id = ?", txId).Update("request_id", requestID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ghr *gameHistoryRepository) CreateGame(gh *models.GameHistory, db *gorm.DB) error {
	return db.Create(gh).Error
}

func NewGameHistoryRepository() models.GameHistoryRepository {
	return &gameHistoryRepository{}
}
