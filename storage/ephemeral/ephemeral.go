package ephemeral

import (
	"fmt"
	"log"
	"split-the-bill-server/storage"
	"split-the-bill-server/types"
	"sync"

	"github.com/google/uuid"
)

type Ephemeral struct {
	lock        sync.Mutex
	userStorage map[uuid.UUID]types.User
	nameIndex   map[string]uuid.UUID
}

func NewEphemeral() *Ephemeral {
	return &Ephemeral{
		userStorage: make(map[uuid.UUID]types.User),
		nameIndex:   make(map[string]uuid.UUID),
	}
}

func (e *Ephemeral) Connect() error {
	return nil
}

func (e *Ephemeral) AddUser(user types.User) (types.User, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if _, ok := e.nameIndex[user.Username]; ok {
		return types.User{}, storage.UserAlreadyExistsError
	}
	_, ok := e.userStorage[user.ID]
	if ok {
		return types.User{}, storage.UserAlreadyExistsError
	}
	// TODO: Test correct usage
	user.ID = uuid.New()
	e.userStorage[user.ID] = user

	e.nameIndex[user.Username] = user.ID
	return user, nil
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

func (e *Ephemeral) CreateAuthCookie(userId uuid.UUID) (types.AuthCookie, error) {
	return types.AuthCookie{}, nil
}

func (e *Ephemeral) LoginUser(types.AuthenticateCredentials) (types.AuthCookie, error) {
	return types.AuthCookie{}, nil
}

func (e *Ephemeral) GetCookieFromUser(userId uuid.UUID) (types.AuthCookie, error) {
	return types.AuthCookie{}, nil
}

func (*Ephemeral) GetUserFromAuthCookie(cookieId uuid.UUID) (types.User, error) {
	return types.User{}, nil
}
