package account

import (
	"challenge/internal/ports"
	"net/smtp"
)

var _ ports.AccountEmailRepository = &RepositoryEmail{}

type RepositoryEmail struct {
	EmailConnect smtp.Auth
	Theme        string
	Logo         string
	From         string
	To           []string
	Subject      string
	Body         string
	SmtpLink     string
}
