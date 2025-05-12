package account

import (
	"challenge/internal/ports"
)

var _ ports.AccountService = &Service{}

type Service struct {
	Account  ports.AccountRepository
	Database ports.AccountDatabaseRepository
	Email    ports.AccountEmailRepository
	Logo     string
}
