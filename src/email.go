package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"

	gomail "gopkg.in/mail.v2"
)

type EmailConfig struct {
	gmail string
	password string
	template string
}

type Email struct {
	from    string
	to      string
	message string
	subject string
}

func (config *EmailConfig) readTemplate(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil { 
		log.Fatal("Unable to read template file")
		return err
	}

	template, err := ioutil.ReadAll(file) 
	if err != nil {
		log.Fatal("Unable to read template file")
		return err
	}
	config.template = string(template)

	return nil
}

func (config *EmailConfig) createEmail(from, to Person) Email {
	
	// replace names in the template
	fromRe := regexp.MustCompile(`{{from}}`)
	toRe := regexp.MustCompile(`{{to}}`)
	message := fromRe.ReplaceAllString(config.template, from.Name)
	message = toRe.ReplaceAllString(message, to.Name)

	return Email {
		from: config.gmail,
		to: from.Email,
		message: message,
		subject: "Fun fun kris kringle time :D",
	}
}

func (config *EmailConfig) send(email Email) error {
	// now we have to send the email
	m := gomail.NewMessage()

	m.SetHeader("From", config.gmail)
	m.SetHeader("To", email.to)
	m.SetHeader("Subject", email.subject)
	m.SetBody("text/html", email.message)

	d := gomail.NewDialer("smtp.gmail.com", 587, config.gmail, config.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil { 
		fmt.Println(err)
		return err
	}
	return nil
}
