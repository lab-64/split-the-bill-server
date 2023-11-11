package authentication

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"split-the-bill-server/core"
	. "split-the-bill-server/domain/model"
	. "split-the-bill-server/storage/storage_inf"
	"time"
)

const SessionCookieValidityPeriod = time.Hour * 24 * 7
const SessionCookieName = "session_cookie"
const ErrMsgAuthentication = "Authentication declined: %v"
const ErrMsgNoCookie = "Authentication cookie is missing"
const ErrMsgInvalidCookie = "Authentication cookie is invalid"

var SessionExpiredError = errors.New("session expired")

type Authenticator struct {
	cookieStorage ICookieStorage
}

func NewAuthenticator(cookieStorage *ICookieStorage) *Authenticator {
	return &Authenticator{cookieStorage: *cookieStorage}
}

// Authenticate checks if the user is authenticated.
// It validates the session cookie, retrieves the authentication cookie, and proceeds if authentication is successful.
func (a *Authenticator) Authenticate(c *fiber.Ctx) error {
	cookie := c.Cookies(SessionCookieName)
	if cookie == "" {
		return core.Error(c, fiber.StatusUnauthorized, ErrMsgNoCookie)
	}

	tokenUUID, err := uuid.Parse(cookie)
	if err != nil {
		return core.Error(c, fiber.StatusUnauthorized, ErrMsgInvalidCookie)
	}

	// get auth cookie from storage
	sessionCookie, err := a.cookieStorage.GetCookieFromToken(tokenUUID)

	if err != nil {
		return core.Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgAuthentication, err))
	}

	err = isSessionCookieValid(sessionCookie)

	if err != nil {
		return core.Error(c, fiber.StatusUnauthorized, fmt.Sprintf(ErrMsgAuthentication, err))
	}

	// go to the next handler
	err = c.Next()

	if err != nil {
		return err
	}

	return err
}

func GenerateSessionCookie(userID uuid.UUID) AuthCookieModel {
	// TODO: Safely generate a session cookie.
	return AuthCookieModel{
		UserID:      userID,
		Token:       uuid.New(),
		ValidBefore: time.Now().Add(SessionCookieValidityPeriod),
	}
}

// isSessionCookieValid validates the given session cookie by checking if the ValidBefore time is in the future. Returns SessionExpiredError if the cookie is not valid anymore.
func isSessionCookieValid(cookie AuthCookieModel) error {
	if cookie.ValidBefore.After(time.Now()) {
		return nil
	} else {
		return SessionExpiredError
	}
}
