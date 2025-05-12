package domain

import (
	"time"
)

type Account struct {
	Id           uint         `gorm:"primaryKey"`
	Name         string       `gorm:"size:255"`
	Total        float32      `gorm:"default:0"`
	Debit        float32      `gorm:"default:0"`
	Credit       float32      `gorm:"default:0"`
	CreatedAt    time.Time    `gorm:"autoCreateTime"`
	Transactions Transactions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Accounts []Account
