package main

import (
	"log"
	"net/mail"
	"time"

	"github.com/scottwcode/go-bnb-bookings/internal/models"
)

func listenForMail() {
	// run func in background so that there will not be a delay in sending mail
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {
	// just creating a dummy mail server
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	// if cannot connect in 10 seconds, just give up
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.April
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
	email.SetBody(mail.TextHTML, "Hello, <strong>world</strong>!")

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent!")
	}

}
