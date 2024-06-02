package dto

type UserRegistration struct {
	UserName     string `json:"user_name" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name"`
	Password     string `json:"password" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	EmailRecheck string `json:"email_recheck" binding:"required,email"`
}

type UserLogin struct {
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	RememberMe bool   `json:"remember_me" binding:"required"`
}
