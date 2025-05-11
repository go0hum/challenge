package models

import (
	"challenge/mysql/database"
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

type Transaction struct {
	Id          uint      `gorm:"primaryKey"`
	Transaction float32   `gorm:"default:0"`
	Date        time.Time `gorm:"autoCreateTime"`
	Account_id  uint      `gorm:"foreignKey:AccountId"`
	Account     Account   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Transactions []Transaction

func Migraciones() {
	database.Database.AutoMigrate(&Account{})
	database.Database.AutoMigrate(&Transaction{})
}
