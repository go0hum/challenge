package handlers

import (
	"challenge/mysql/database"
	"challenge/mysql/models"
	"fmt"
)

func GetAccounts() {
	accounts := models.Accounts{}
	err := database.Database.Order("id desc").Find(&accounts).Error
	if err != nil {
		panic(err)
	}
	for _, account := range accounts {
		fmt.Println(account.Id, account.Total, account.Debit, account.Credit)
	}
}

func SetAccount(name string, total float32, debit float32, credit float32) {
	account := models.Account{Name: name, Total: total, Debit: debit, Credit: credit}
	err := database.Database.Create(&account).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("Account created:", account.Id, account.Total, account.Debit, account.Credit)
}

func SetAccounts(account models.Account) {
	result := database.Database.Create(&account)
	if result.Error != nil {
		fmt.Println("Account failed:", result.Error)
	} else {
		fmt.Println("Account succeeded")
	}
}
