package main_test

import (
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"split-the-bill-server/authentication"
	"split-the-bill-server/domain/service/impl"
	"split-the-bill-server/presentation/handler"
	"split-the-bill-server/presentation/router"
	"split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/ephemeral/eph_storages"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// TestlandingPage tests whether the fiber client starts correctly.
func TestLandingPage(t *testing.T) {
	// Test Server Configuration
	app := fiber.New()
	e, err := ephemeral.NewEphemeral()
	require.NoError(t, err)
	v, err := authentication.NewPasswordValidator()
	require.NoError(t, err)

	//storages
	userStorage := eph_storages.NewUserStorage(e)
	groupStorage := eph_storages.NewGroupStorage(e)
	cookieStorage := eph_storages.NewCookieStorage(e)
	billStorage := eph_storages.NewBillStorage(e)
	invitationStorage := eph_storages.NewInvitationStorage(e)

	//services
	userService := impl.NewUserService(&userStorage, &cookieStorage)
	groupService := impl.NewGroupService(&groupStorage, &userStorage)
	billService := impl.NewBillService(&billStorage, &groupStorage)
	invitationService := impl.NewInvitationService(&invitationStorage, &groupStorage)

	//handlers
	userHandler := handler.NewUserHandler(&userService, v)
	groupHandler := handler.NewGroupHandler(&groupService, &invitationService)
	billHandler := handler.NewBillHandler(&billService, &groupService)
	invitationHandler := handler.NewInvitationHandler(&invitationService)

	// authenticator
	authenticator := authentication.NewAuthenticator(&cookieStorage)

	//routing
	router.SetupRoutes(app, *userHandler, *groupHandler, *billHandler, *invitationHandler, *authenticator)

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
