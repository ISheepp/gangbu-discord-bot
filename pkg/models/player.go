package models

import (
	"gorm.io/gorm"
	"time"
)

// Player struct gamer
type Player struct {
	gorm.Model
	DiscordUserId string `gorm:"unique;not null"` // discord user id
	WalletAddress string // wallet address
	PrivateKey    string // private key
	WalletValue   int64  // wallet value 「unit gwei」
}

type PlayerVo struct {
	ID                 uint
	DiscordUserId      string
	WalletAddress      string
	WalletValue        int64
	WithDrawHistoryDto []WithDrawHistoryDto
	CreateAt           time.Time
}

type PlayerCreateBo struct {
	DiscordUserId string
}

type PlayerRepository interface {
	GetByDiscordUserID(discordUserID string) (*Player, error)
	CreatePlayer(p *Player) error
	UpdateWalletValue(id uint, walletValue int64) error
}

type PlayerUsecase interface {
	GetByDiscordUserIDOrCreate(discordUserID string) (*PlayerVo, error)
	CreatePlayer(bo PlayerCreateBo) error
	GetByDiscordUserID(discordUserID string) (*Player, error)
}
