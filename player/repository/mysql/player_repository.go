package mysql

import (
	"gangbu/pkg/models"
	"gorm.io/gorm"
)

type playerRepository struct {
	db *gorm.DB
}

func (pr *playerRepository) UpdateWalletValue(playerId uint, walletValue int64) error {
	return pr.db.Model(&models.Player{}).Where("id = ?", playerId).Update("wallet_value", walletValue).Error
}

func (pr *playerRepository) GetByDiscordUserID(discordUserID string) (*models.Player, error) {
	player := &models.Player{}
	if e := pr.db.Where("discord_user_id = ?", discordUserID).First(player).Error; e != nil {
		return nil, e
	}
	return player, nil
}

func (pr *playerRepository) CreatePlayer(p *models.Player) error {
	if e := pr.db.Create(p).Error; e != nil {
		return e
	}
	return nil
}

func NewPlayerRepository(db *gorm.DB) models.PlayerRepository {
	return &playerRepository{
		db: db,
	}

}
