package main

import (
	"challenge/csv/file"
	"challenge/email"
	"challenge/mysql/handlers"
	"challenge/mysql/models"
	"fmt"

	"github.com/jinzhu/copier"
)

func main() {
	accountReadFile := file.ReadFile()

	var transactions []models.Transaction
	copier.Copy(&transactions, &accountReadFile.Transactions)

	account := models.Account{
		Name:         accountReadFile.Name,
		Total:        accountReadFile.Total,
		Debit:        accountReadFile.Debit,
		Credit:       accountReadFile.Credit,
		Transactions: transactions,
	}
	handlers.SetAccounts(account)

	var conteoMes map[string]int8
	copier.Copy(&conteoMes, &accountReadFile.Months)

	templateData := struct {
		Logo        string
		Total       float32
		TotalDebit  float32
		TotalCredit float32
		ConteoMes   map[string]int8
	}{
		Logo:        "https://challenge-storicard.s3.us-east-1.amazonaws.com/logo.png",
		Total:       accountReadFile.Total,
		TotalDebit:  accountReadFile.Debit,
		TotalCredit: accountReadFile.Credit,
		ConteoMes:   conteoMes,
	}
	email.ConnectSMTP()
	r := email.NewRequest("zooxial@gmail.com", []string{"1201hs@gmail.com"}, "Balance General", "")
	if err := r.ParseTemplate("../../email/template.html", templateData); err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	} else {
		fmt.Println("Error parsing template:", err)
	}
}
