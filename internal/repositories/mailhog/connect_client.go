package mailhog

import (
	"net/smtp"
	"strings"
)

func ConnectClient(email string, password string, smtpLink string) smtp.Auth {
	parts := strings.Split(smtpLink, ":")
	serverSmtp := parts[0]
	return smtp.PlainAuth("", email, password, serverSmtp)
}
