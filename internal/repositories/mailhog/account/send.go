package account

import (
	"bytes"
	domain "challenge/internal/domain/account"
	"html/template"
	"log"
	"net/smtp"
)

func (re *RepositoryEmail) ParseTemplate(data domain.TemplateEmail) error {
	t, err := template.ParseFiles(re.Theme)
	if err != nil {
		log.Fatal("error template: ", err)
		return err
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		log.Fatal("Error sending email:", err)
		return err
	}
	re.Body = buf.String()
	return nil
}

func (re *RepositoryEmail) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + re.Subject + "!\n"
	msg := []byte(subject + mime + "\n" + re.Body)
	var auth smtp.Auth
	if re.SmtpLink == "localhost:1025" {
		auth = nil
	} else {
		auth = re.EmailConnect
	}
	addr := re.SmtpLink

	if err := smtp.SendMail(addr, auth, re.From, re.To, msg); err != nil {
		log.Fatal("Error sending email:", err)
		return false, err
	}
	log.Println("Correo enviado")
	return true, nil
}
