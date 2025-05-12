package account

import (
	"challenge/internal/ports"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var _ ports.AccountRepository = &Repository{}

type Repository struct {
	Client *s3.Client
	Bucket string
	Key    string
}
