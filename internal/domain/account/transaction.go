package domain

import "time"

type Transaction struct {
	Id          uint      `gorm:"primaryKey"`
	Transaction float32   `gorm:"default:0"`
	Date        time.Time `gorm:"autoCreateTime"`
	Account_id  uint      `gorm:"foreignKey:AccountId"`
	Account     Account   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Transactions []Transaction
