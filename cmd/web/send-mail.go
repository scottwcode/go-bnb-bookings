package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/scottwcode/go-bnb-bookings/internal/models"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}() // run func in background
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	// if cannot connect in 10 seconds, just give up
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	// for a production mail server will need the following
	// server.Username
	// server.Password
	// server.Encryption
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
			errorLog.Println(err)
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
