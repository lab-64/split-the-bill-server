package ephemeral

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"split-the-bill-server/types"
	"sync"
)

type Ephemeral struct {
	lock        sync.Mutex
	userStorage map[uuid.UUID]types.User
	nameIndex   map[string]uuid.UUID
}

func NewEphemeral() *Ephemeral {
	return &Ephemeral{
		userStorage: make(map[uuid.UUID]types.User),
	}
}

func (e *Ephemeral) Connect() error {
	return nil
}

func (e *Ephemeral) AddUser(user types.User) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	if _, ok := e.nameIndex[user.Username]; ok {
		return fmt.Errorf("user with username %s already exists", user.Username)
	}
	_, ok := e.userStorage[user.ID]
	if ok {
		log.Println(fmt.Sprintf("FATAL error: user with id %v already exists", user.ID))
		return fmt.Errorf("user with id %v already exists", user.ID)
	}
	e.userStorage[user.ID] = user

	e.nameIndex[user.Username] = user.ID
	return nil
}

func (e *Ephemeral) DeleteUser(id uuid.UUID) error {
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.userStorage, id)
	delete(e.nameIndex, e.userStorage[id].Username)
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
		return user, fmt.Errorf("no user with id %v", id)
	}
	return user, nil
}

func (e *Ephemeral) GetUserByUsername(username string) (types.User, error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	id, ok := e.nameIndex[username]
	if !ok {
		return types.User{}, fmt.Errorf("no user with username %s", username)
	}
	user, ok := e.userStorage[id]
	if !ok {
		log.Println(fmt.Sprintf("FATAL error: user storage inconsistent: username '%s' points to non-existent user", username))
		return user, fmt.Errorf("user storage inconsistent: username '%s' points to non-existent user", username)
	}
	return user, nil
}
