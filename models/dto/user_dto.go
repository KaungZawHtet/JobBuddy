package dto

import "github.com/google/uuid"

type UserRegistration struct {
	UserName     string `json:"user_name" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name"`
	Password     string `json:"password" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	EmailRecheck string `json:"email_recheck" binding:"required,email"`
}

type UserLogin struct {
	Id         uuid.UUID
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	RememberMe bool   `json:"remember_me" binding:"required"`
}

type GoogleTokenResponse struct {
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
}

type UserProfile struct {
	Email         string `json:"email"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	ID            string `json:"id"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
