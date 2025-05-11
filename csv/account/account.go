package account

import "time"

type TransactionFile struct {
	Transaction float32
	Date        time.Time
}

type MonthsFile map[string]int8

type AccountFile struct {
	Name         string
	Total        float32
	Debit        float32
	Credit       float32
	Transactions []TransactionFile
	Months       MonthsFile
}
