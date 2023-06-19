package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dapetoo/go-bookings/internal/models"
	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMessage(msg)
		}
	}()
}

func sendMessage(m models.MailData) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := mail.NewSMTPClient()
	fmt.Println(os.Getenv("MAIL_SERVER"))
	log.Println(os.Getenv("MAIL_SERVER"))
	log.Println(os.Getenv("MAIL_USERNAME"))
	log.Println(os.Getenv("MAIL_PASSWORD"))
	server.Host = os.Getenv("MAIL_SERVER")
	server.Port = 2525
	server.Username = os.Getenv("MAIL_USERNAME")
	server.Password = os.Getenv("MAIL_PASSWORD")
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	if m.Template == "" {
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		data, err := os.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
		}

		mailTemplate := string(data)
		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)
		email.SetBody(mail.TextHTML, msgToSend)
	}
	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent!")
	}
}
