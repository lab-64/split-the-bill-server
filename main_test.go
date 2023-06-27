package main_test

import (
	"net/http/httptest"
	"split-the-bill-server/handler"
	"split-the-bill-server/router"
	"split-the-bill-server/storage/ephemeral"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestLandingPage test the configuration of the default landing page
//
// Parameters:
//
//	*testing.T: Testing type
func TestLandingPage(t *testing.T) {
	// Test Server Configuration
	app := fiber.New()
	storage := ephemeral.NewEphemeral()
	h := handler.NewHandler(storage)
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
