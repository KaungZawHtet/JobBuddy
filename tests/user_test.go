// main_test.go

package tests

import (
	"JobBuddy/models/domain"
	"JobBuddy/services"
	"JobBuddy/types"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"

	"JobBuddy/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/ping", handlers.HandlePing)

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"message":"pong"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetUser(t *testing.T) {

	services.CreateUser(&domain.User{
		Email:                  "kaungzawhtet.mm@gmail.com",
		UserName:               "KZH",
		Password:               "f23fjqvnqv2420tt14",
		EmailConfirmed:         true,
		EmailConfirmationToken: "ef20f4-qn4v4q3vjq3jvqj",
	})

	result, _ := services.GetUser(types.ByEmail, "kaungzawhtet.mm@gmail.com")

	assert.Equal(t, result.Email, "kaungzawhtet.mm@gmail.com")

}
