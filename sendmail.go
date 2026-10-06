package sscraper

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendMail(text_mail string,subject string,sender string, reciver string, token string) error {
	message := gomail.NewMessage()

	// Using Mailtrap's demo domain (no verification needed!)
	message.SetHeader("From", sender)
	message.SetHeader("To", reciver) // Your real email
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", text_mail)

	// Your actual Mailtrap credentials
	dialer := gomail.NewDialer(
		"live.smtp.mailtrap.io", // or "smtp.mailtrap.io"
		587,
		"api",                              // Username is always "api"
		"######################", // Replace with your token from Mailtrap
	)

	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
		return err
	} else {
		fmt.Println(" \nEmail sent successfully to alishadan84@gmail.com!")
	}
	return nil
}
