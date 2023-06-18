package main

import (
	"fmt"
	"github.com/dapetoo/go-bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
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
	server := mail.NewSMTPClient()
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
		data, err := ioutil.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
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
