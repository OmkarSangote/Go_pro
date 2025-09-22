// 4. Sending Emails
// Example: Automating Email Notifications

// You can automate sending emails using Go's net/smtp package.

package main

import (
	"log"
	"net/smtp"
)

func main() {
	from := "madhushruti@outlook.com"
	password := "maakalidevi18" // Use an app password if 2FA is enabled
	to := []string{"madhusmriti@outlook.com"}
	smtpHost := "smtp.office365.com"
	smtpPort := "587"

	msg := []byte("To: madhusmriti@outlook.com\r\n" +
		"Subject: Test Email\r\n" +
		"\r\n" +
		"This is a test email sent from Go.\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Email sent successfully!")
}
