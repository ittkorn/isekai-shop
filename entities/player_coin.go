package entities

import "time"

type PlayerCoin struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	PlayerID  string    `gorm:"type:varchar(64);not null"`
	Amount    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"` // หรือ `gorm:""` หากต้องการ default gorm timestamp
}