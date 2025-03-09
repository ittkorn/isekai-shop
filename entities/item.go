package entities

import "time"

type Item struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	AdminID     string    `gorm:"type:varchar(64)"`
	Name        string    `gorm:"type:varchar(64);unique;not null"`
	Description string    `gorm:"type:varchar(64);unique;not null"`
	Picture     string    `gorm:"type:varchar(64);not null"` // หรือ `gorm:"type:varchar(64);unique;not null"` หรือ `gorm:"type:varchar(64)"` แล้วแต่ requirement
	Price       uint      `gorm:"not null"`
	Archive     bool      `gorm:"default:false"` // หรือ `gorm:"not null;default:false"` แล้วแต่ requirement
	CreatedAt   time.Time `gorm:""`             // หรือ `gorm:"not null"` แล้วแต่ requirement
	UpdatedAt   time.Time
}