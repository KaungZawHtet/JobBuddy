package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

func SendEmailConfirmationLink(email, emailToken string) (bool, error) {

	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	sender := fmt.Sprintf("JobBuddy<contact@%s>", domain)

	mg := mailgun.NewMailgun(domain, apiKey)

	subject := "Email Confirmation"

	message := mg.NewMessage(sender, subject, "", email)
	body := fmt.Sprintf(`
<html>
<body>
	<h1>Hello, Please confirm your email</h1>
	<a href="http://localhost:8080/api/user/email-confirm">Here </a>
</body>
</html>
`)

	message.SetHtml(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	_, _, errEmail := mg.Send(ctx, message) // resp, id can be checked

	if errEmail != nil {

		println(errEmail.Error())

		return false, errEmail

	}

	/*
		NOTE: for debugging
		println("This is resp")
		println(resp)
		println("This is id")
		println(id) */

	return true, nil

}

func GenerateEmailToken() (string, error) {
	// Generate random bytes
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	// Encode random bytes to base64
	token := base64.URLEncoding.EncodeToString(randomBytes)
	return token, nil
}
