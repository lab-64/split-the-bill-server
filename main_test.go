package main

import (
	"net/http/httptest"
	"split-the-bill-server/server"
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
	server.SetupRoutes(app)

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
