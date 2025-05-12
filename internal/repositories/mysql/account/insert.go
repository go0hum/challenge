package account

import (
	domain "challenge/internal/domain/account"
	"log"
)

func (r *RepositoryMysql) Insert(account domain.Account) error {
	result := r.Db.Create(&account)
	if result.Error != nil {
		log.Fatal("Account failed:", result.Error)
		return result.Error
	}
	log.Println("Registro insertado correctamente")
	return nil
}
