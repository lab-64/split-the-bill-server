package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"split-the-bill-server/core"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
	. "split-the-bill-server/storage/database/test_util"
	storagemocks "split-the-bill-server/storage/mocks"
	"testing"
	"time"
)

var (
	app           *fiber.App
	authenticator *Authenticator
)

func TestMain(m *testing.M) {
	// setup authentication
	cookieStorage := storagemocks.NewCookieStorageMock()
	authenticator = NewAuthenticator(&cookieStorage)

	// setup fiber
	app = fiber.New()
	// setup test route
	app.Get("/user", authenticator.Authenticate, func(c *fiber.Ctx) error { return core.Success(c, fiber.StatusOK, "Authentication accept", nil) })

	// Run tests
	exitCode := m.Run()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}

func TestAuthenticate_Success(t *testing.T) {

	// mock get cookie from storage
	storagemocks.MockCookieGetCookieFromToken = func(token uuid.UUID) (model.AuthCookieModel, error) {
		return UserCookie, nil
	}

	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	// add cookie to request
	req.AddCookie(&http.Cookie{
		Name:  "session_cookie",
		Value: UserCookie.Token.String(),
	})
	// perform request
	resp, err := app.Test(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Fatal("Error in test setup during performing a request - ", err)
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestAuthenticate_NoCookie(t *testing.T) {
	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	// perform request
	resp, err := app.Test(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Fatal("Error in test setup during performing a request - ", err)
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error in test setup during reading response body - ", err)
	}
	var response dto.GeneralResponseDTO
	err = json.Unmarshal(body, &response)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, ErrMsgNoCookie, response.Message)
}

func TestAuthenticate_InvalidCookie(t *testing.T) {
	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	// add cookie to request
	req.AddCookie(&http.Cookie{
		Name:  "session_cookie",
		Value: "invalid",
	})
	// perform request
	resp, err := app.Test(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Fatal("Error in test setup during performing a request - ", err)
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error in test setup during reading response body - ", err)
	}
	var response dto.GeneralResponseDTO
	err = json.Unmarshal(body, &response)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, ErrMsgInvalidCookie, response.Message)
}

func TestAuthenticate_DeclineCookie(t *testing.T) {

	// mock get cookie from storage
	storagemocks.MockCookieGetCookieFromToken = func(token uuid.UUID) (model.AuthCookieModel, error) {
		return model.AuthCookieModel{}, storage.NoSuchCookieError
	}

	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	// add cookie to request
	req.AddCookie(&http.Cookie{
		Name:  "session_cookie",
		Value: UserCookie.Token.String(),
	})
	// perform request
	resp, err := app.Test(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Fatal("Error in test setup during performing a request - ", err)
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error in test setup during reading response body - ", err)
	}
	var response dto.GeneralResponseDTO
	err = json.Unmarshal(body, &response)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, fmt.Sprintf(ErrMsgAuthentication, storage.NoSuchCookieError), response.Message)

}

func TestAuthenticate_ExpiredCookie(t *testing.T) {

	// mock get cookie from storage
	storagemocks.MockCookieGetCookieFromToken = func(token uuid.UUID) (model.AuthCookieModel, error) {
		return model.AuthCookieModel{
			UserID:      User.ID,
			Token:       token,
			ValidBefore: time.Now().Add(-time.Hour),
		}, nil
	}

	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	// add cookie to request
	req.AddCookie(&http.Cookie{
		Name:  "session_cookie",
		Value: UserCookie.Token.String(),
	})
	// perform request
	resp, err := app.Test(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Fatal("Error in test setup during performing a request - ", err)
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error in test setup during reading response body - ", err)
	}
	var response dto.GeneralResponseDTO
	err = json.Unmarshal(body, &response)

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, fmt.Sprintf(ErrMsgAuthentication, SessionExpiredError), response.Message)
}
