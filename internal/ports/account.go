package ports

import (
	domain "challenge/internal/domain/account"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AccountService interface {
	Get(ctx context.Context) (domain.AccountFile, error)
	Create(account domain.AccountFile)
	Send(accountReadFile domain.AccountFile)
}

type AccountRepository interface {
	ReadCSV(ctx context.Context) (*s3.GetObjectOutput, error)
}

type AccountDatabaseRepository interface {
	Insert(account domain.Account) error
}

type AccountEmailRepository interface {
	ParseTemplate(templateEmail domain.TemplateEmail) error
	SendEmail() (bool, error)
}
