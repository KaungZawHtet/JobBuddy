package services

import (
	"crypto/rand"
	"encoding/base64"
)

func sendEmailConfirmationLink(email, emailToken string) {

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
