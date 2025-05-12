package account

import (
	domain "challenge/internal/domain/account"

	"github.com/jinzhu/copier"
)

func (s *Service) Create(accountReadFile domain.AccountFile) {
	var transactions []domain.Transaction
	copier.Copy(&transactions, &accountReadFile.Transactions)

	account := domain.Account{
		Name:         accountReadFile.Name,
		Total:        accountReadFile.Total,
		Debit:        accountReadFile.Debit,
		Credit:       accountReadFile.Credit,
		Transactions: transactions,
	}
	s.Database.Insert(account)
}
