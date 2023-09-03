package models

import (
	"gorm.io/gorm"
	"time"
)

type WithDrawHistory struct {
	gorm.Model
	PlayerDiscordUserId string
	WithdrawValue       int64
	WithDrawTime        time.Time
	WithDrawAddress     string
	WithdrawStatus      int
	ServerId            string // discord server id
}
