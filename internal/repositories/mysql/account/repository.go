package account

import (
	"challenge/internal/ports"

	"gorm.io/gorm"
)

var _ ports.AccountDatabaseRepository = &RepositoryMysql{}

type RepositoryMysql struct {
	Db *gorm.DB
}
