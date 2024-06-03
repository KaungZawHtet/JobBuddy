// main_test.go

package tests

import (
	"JobBuddy/services"
	"JobBuddy/types"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Load the .env file
	godotenv.Load()

	// Run the tests
	code := m.Run()

	// Exit with the test code
	os.Exit(code)
}

func TestPingRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	expectedBody := `{"message":"pong"}`
	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestGetUser(t *testing.T) {

	result, _ := services.GetUser(types.ByEmail, "kaungzawhtet.mm@gmail.com")

	assert.Equal(t, result.Email, "kaungzawhtet.mm@gmail.com")

}
