package ephemeral

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
	"sync"
)

type Ephemeral struct {
	lock            sync.Mutex
	userStorage     map[uuid.UUID]types.User
	nameIndex       map[string]uuid.UUID
	passwordStorage map[uuid.UUID][]byte
	cookieStorage   map[uuid.UUID][]types.AuthenticationCookie
}

func (e *Ephemeral) AddAuthenticationCookie(cookie types.AuthenticationCookie) {
	e.lock.Lock()
	defer e.lock.Unlock()
	cookies, exists := e.cookieStorage[cookie.UserID]
	if !exists {
		cookies = make([]types.AuthenticationCookie, 0)
	}
	cookies = append(cookies, cookie)
	e.cookieStorage[cookie.UserID] = cookies
}

func (e *Ephemeral) GetCookiesForUser(userID uuid.UUID) []types.AuthenticationCookie {
	e.lock.Lock()
	defer e.lock.Unlock()
	return e.cookieStorage[userID]
}

func (e *Ephemeral) GetCredentials(id uuid.UUID) ([]byte, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	hash, exists := e.passwordStorage[id]
	if !exists {
		return nil, storage.NoCredentialsError
	}
	return hash, nil
}

func (e *Ephemeral) RegisterUser(user types.User, hash []byte) error {
	err := e.AddUser(user)
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	_, exists := e.passwordStorage[user.ID]
	if exists {
		return errors.New("fatal: user already has saved password")
	}
	e.passwordStorage[user.ID] = hash
	return nil
}

func NewEphemeral() *Ephemeral {
	return &Ephemeral{
		userStorage:     make(map[uuid.UUID]types.User),
		nameIndex:       make(map[string]uuid.UUID),
		passwordStorage: make(map[uuid.UUID][]byte),
		cookieStorage:   make(map[uuid.UUID][]types.AuthenticationCookie),
	}
}

func (e *Ephemeral) Connect() error {
	return nil
}

func (e *Ephemeral) AddUser(user types.User) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	if _, ok := e.nameIndex[user.Username]; ok {
		return storage.UserAlreadyExistsError
	}
	_, ok := e.userStorage[user.ID]
	if ok {
		return storage.UserAlreadyExistsError
	}
	e.userStorage[user.ID] = user
	e.nameIndex[user.Username] = user.ID
	return nil
}

func (e *Ephemeral) DeleteUser(id uuid.UUID) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	user, exists := e.userStorage[id]
	if !exists {
		return nil
	}
	delete(e.nameIndex, user.Username)
	delete(e.userStorage, id)
	delete(e.passwordStorage, id)
	return nil
}

func (e *Ephemeral) GetAllUsers() ([]types.User, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	users := make([]types.User, len(e.userStorage))
	i := 0
	for _, user := range e.userStorage {
		users[i] = user
		i++
	}
	return users, nil
}

func (e *Ephemeral) GetUserByID(id uuid.UUID) (types.User, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	user, ok := e.userStorage[id]
	if !ok {
		return user, storage.NoSuchUserError
	}
	return user, nil
}

func (e *Ephemeral) GetUserByUsername(username string) (types.User, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	id, ok := e.nameIndex[username]
	if !ok {
		return types.User{}, storage.NoSuchUserError
	}
	user, ok := e.userStorage[id]
	if !ok {
		log.Printf("FATAL error: user storage inconsistent: username '%s' points to non-existent user", username)
		return user, fmt.Errorf("user storage inconsistent: username '%s' points to non-existent user", username)
	}
	return user, nil
}

func (e *Ephemeral) GetCookieFromToken(token uuid.UUID) (types.AuthenticationCookie, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	for _, cookies := range e.cookieStorage {
		for _, cookie := range cookies {
			if cookie.Token == token {
				return cookie, nil
			}
		}
	}
	return types.AuthenticationCookie{}, storage.NoSuchCookieError
}
