// tests/main_test.go

package tests

import (
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Load the .env file

	setup()
	godotenv.Load()

	// Run the tests
	code := m.Run()

	// Teardown
	teardown()

	// Exit with the test code
	os.Exit(code)
}

func setup() {
	// Your setup logic here
}

func teardown() {
	// Your teardown logic here
}
