package main_test

import (
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"split-the-bill-server/authentication"
	"split-the-bill-server/handler"
	"split-the-bill-server/router"
	"split-the-bill-server/storage/ephemeral"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestlandingPage tests whether the fiber client starts correctly.
func TestLandingPage(t *testing.T) {
	// Test Server Configuration
	app := fiber.New()
	storage := ephemeral.NewEphemeral()
	v, err := authentication.NewPasswordValidator()
	require.NoError(t, err)
	h := handler.NewHandler(storage, storage, storage, v)
	router.SetupRoutes(app, h)

	// Create a new http get request on landingpage
	req := httptest.NewRequest("GET", "/", nil)

	// Perform request
	resp, _ := app.Test(req, 1)

	// Testing logs
	t.Log(resp.StatusCode)
	t.Log(resp)
	// Expect HTTP.OK
	assert.Equal(t, 200, resp.StatusCode)

}
