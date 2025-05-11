package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

var auth smtp.Auth

func ConnectSMTP() {
	auth = smtp.PlainAuth("", "**CORREO**", "**PASSWORD**", "smtp.gmail.com")
}

type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(from string, to []string, subject, body string) *Request {
	return &Request{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "localhost:1025" //"smtp.gmail.com:587"

	if err := smtp.SendMail(addr, nil, r.from, r.to, msg); err != nil {
		fmt.Println("Error sending email:", err)
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()

	return nil
}
