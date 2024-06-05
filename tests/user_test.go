// main_test.go

package tests

import (
	"JobBuddy/models/domain"
	"JobBuddy/services"
	"JobBuddy/types"

	"github.com/stretchr/testify/assert"
	"os"

	"testing"
)

func TestGetUser(t *testing.T) {

	env := os.Getenv("ENVIRONMENT")

	if env == "test" {

		services.CreateUser(&domain.User{
			Email:                  "kaungzawhtet.mm@gmail.com",
			UserName:               "KZH",
			Password:               "f23fjqvnqv2420tt14",
			EmailConfirmed:         true,
			EmailConfirmationToken: "ef20f4-qn4v4q3vjq3jvqj",
		})

	}

	result, _ := services.GetUser(types.ByEmail, "kaungzawhtet.mm@gmail.com")

	assert.Equal(t, result.Email, "kaungzawhtet.mm@gmail.com")

}
