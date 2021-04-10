package implementations

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
)

type IMailProvider interface {
	SendMail(email, name, language, subject string)
}

type mailRepository struct {
	test string
}

func NewMailRepository() IMailProvider {
	return &mailRepository{"Gabriel"}
}

func (mailRepository *mailRepository) SendMail(email, name, language, subject string) {
	auth := smtp.PlainAuth("", "ab9923e8f2fbb7", "c667f912b0b652", "smtp.mailtrap.io")
	from := "not-reply@ownershop.com"

	templateData := struct {
		UserName string
	}{
		UserName: name,
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	to := []string{email}
	body := mailRepository.parseTemplate("./src/templates/email/create_account.html", templateData)
	subjectBody := "Subject: " + subject + "!\n"
	msg := []byte(subjectBody + mime + "\n" + body)

	// Check why from and to doesnt fill properly
	err := smtp.SendMail("smtp.mailtrap.io:2525", auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func (mailRepository *mailRepository) parseTemplate(templateFileName string, data interface{}) string {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		fmt.Println(err)
	}
	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		fmt.Println(err)
	}

	return buf.String()
}
