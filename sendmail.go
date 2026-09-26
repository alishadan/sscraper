package sscraper

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendMail(extracted_price string, url string, product string) error {
	string_all := "Hi Ali,\nprice of " + product + " in " + url +
		" has more than 5% change\n" +
		"now price of A4 paper double A in this site is: " + extracted_price +
		"\n"

	message := gomail.NewMessage()

	// Using Mailtrap's demo domain (no verification needed!)
	message.SetHeader("From", "hello@demomailtrap.co")
	message.SetHeader("To", "alishadan84@gmail.com") // Your real email
	message.SetHeader("Subject", "Price Change Alert")
	message.SetBody("text/plain", string_all)

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
