package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

var auth smtp.Auth

func main() {
	auth = smtp.PlainAuth("", "**EMAIL**", "**PASSWORD**", "smtp.gmail.com")
	templateData := struct {
		Name        string
		URL         string
		Logo        string
		Total       float32
		TotalDebit  float32
		TotalCredit float32
		ConteoMes   map[string]int8
	}{
		Name:        "Humberto Santos",
		URL:         "https://humberto.zooxial.com",
		Logo:        "https://challenge-storicard.s3.us-east-1.amazonaws.com/logo.png",
		Total:       39.74,
		TotalDebit:  -15.38,
		TotalCredit: 35.25,
		ConteoMes: map[string]int8{
			"Enero":   1,
			"Febrero": 2,
		},
	}
	r := NewRequest("zooxial@gmail.com", []string{"1201hs@gmail.com"}, "Balance General", "")
	if err := r.ParseTemplate("template.html", templateData); err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}
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
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, r.from, r.to, msg); err != nil {
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
