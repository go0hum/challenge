package account

import (
	domain "challenge/internal/domain/account"
	"fmt"

	"github.com/jinzhu/copier"
)

func (s *Service) Send(accountReadFile domain.AccountFile) {
	var conteoMes map[string]int8
	copier.Copy(&conteoMes, &accountReadFile.Months)

	templateEmail := domain.TemplateEmail{
		Logo:        s.Logo,
		Total:       accountReadFile.Total,
		TotalDebit:  accountReadFile.Debit,
		TotalCredit: accountReadFile.Credit,
		ConteoMes:   conteoMes,
	}

	if err := s.Email.ParseTemplate(templateEmail); err == nil {
		s.Email.SendEmail()
	} else {
		fmt.Println("Error parsing template:", err)
	}
}
