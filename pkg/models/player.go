package models

import (
	"gorm.io/gorm"
)

// Player struct gamer
type Player struct {
	gorm.Model
	DiscordUserId string `gorm:"unique;not null"` // discord user id
	WalletAddress string // wallet address
	PrivateKey    string // private key
	WalletValue   int64  // wallet value 「unit gwei」
}
