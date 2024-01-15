package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/storage"
	storagemocks "split-the-bill-server/storage/mocks"
	"testing"
	"time"
)

// Testdata
var (
	TestUser = model.UserModel{
		ID:    uuid.New(),
		Email: "test@mail.com",
	}
	TestUserCookie = model.AuthCookieModel{
		UserID:      TestUser.ID,
		Token:       uuid.New(),
		ValidBefore: time.Now().Add(SessionCookieValidityPeriod),
	}
)

func TestAuthenticate_Success(t *testing.T) {

	// mock get cookie from storage
	storagemocks.MockCookieGetCookieFromToken = func(token uuid.UUID) (model.AuthCookieModel, error) {
		return TestUserCookie, nil
	}

	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	// add cookie to request
	req.AddCookie(&http.Cookie{
		Name:  "session_cookie",
		Value: TestUserCookie.Token.String(),
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

	assert.NoError(t, err)
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

	assert.NoError(t, err)
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
		Value: TestUserCookie.Token.String(),
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

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, fmt.Sprintf(ErrMsgAuthentication, storage.NoSuchCookieError), response.Message)

}

func TestAuthenticate_ExpiredCookie(t *testing.T) {

	// mock get cookie from storage
	storagemocks.MockCookieGetCookieFromToken = func(token uuid.UUID) (model.AuthCookieModel, error) {
		return model.AuthCookieModel{
			UserID:      TestUser.ID,
			Token:       token,
			ValidBefore: time.Now().Add(-time.Hour),
		}, nil
	}

	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	// add cookie to request
	req.AddCookie(&http.Cookie{
		Name:  "session_cookie",
		Value: TestUserCookie.Token.String(),
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

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, fmt.Sprintf(ErrMsgAuthentication, SessionExpiredError), response.Message)
}
