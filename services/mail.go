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
	body := `
<html>
<body>
	<h1>Sending HTML emails with Mailgun</h1>
	<p style="color:blue; font-size:30px;">Hello world</p>
	<p style="font-size:30px;">More examples can be found <a href="https://documentation.mailgun.com/en/latest/api-sending.html#examples">here</a></p>
</body>
</html>
`

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
