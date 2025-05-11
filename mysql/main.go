package main

import (
	"challenge/mysql/handlers"
	"challenge/mysql/models"
	"time"
)

func main() {
	//models.Migraciones()
	date1, _ := time.Parse("2006-01-02", "2025-01-15")
	date2, _ := time.Parse("2006-01-02", "2025-02-15")

	transactions := models.Transactions{
		{
			Transaction: 100,
			Date:        date1,
		},
		{
			Transaction: 100,
			Date:        date2,
		},
	}

	account := models.Account{
		Name:         "excel1.csv",
		Total:        200,
		Debit:        -100,
		Credit:       100,
		Transactions: transactions,
	}
	handlers.SetAccounts(account)
}
