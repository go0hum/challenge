package handlers

import (
	"challenge/mysql/database"
	"challenge/mysql/models"
	"fmt"
)

func SetTransactions(transactions []*models.Transaction) {
	result := database.Database.Create(&transactions)
	if result.Error != nil {
		fmt.Println("Transaction failed:", result.Error)
	} else {
		fmt.Println("Transaction succeeded")
	}
}
