package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rohanrj3296/Bread_Butter_Bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}

	}()
}

/*

server := mail.NewSMTPClient()
server.Host = "smtp.gmail.com"
server.Port = 587 // TLS/STARTTLS port
server.Username = "your-gmail-username@gmail.com"
server.Password = "your-app-password" // App-specific password
server.Encryption = mail.EncryptionSTARTTLS



*/

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = "rohanjadhav32963296@gmail.com"
	server.Password = "kxkl pjkw jful zycz"
	server.Encryption = mail.EncryptionSTARTTLS
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	email := mail.NewMSG() // create an empty message
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)

	if m.Template == "" {
		// Set body directly if no template is used
		email.SetBody(mail.TextHTML, m.Content)
	} else {
		// Read the template file
		data, err := os.ReadFile(fmt.Sprintf("./email-templates/%s", m.Template))
		if err != nil {
			app.ErrorLog.Println(err)
			return
		}

		// Replace placeholder in template with actual content
		mailTemplate := string(data)
		msgToSend := strings.Replace(mailTemplate, "[%body%]", m.Content, 1)

		// Set the email body with the final content
		email.SetBody(mail.TextHTML, msgToSend)
	}

	// Try sending the email
	err = email.Send(client)
	if err != nil {
		app.ErrorLog.Println(err)
	} else {
		app.InfoLog.Println("Email sent successfully!")
	}
}
