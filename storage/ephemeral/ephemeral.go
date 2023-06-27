package ephemeral

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
)

type Ephemeral struct {
	userStorage map[uuid.UUID]types.User
}

func NewEphemeral() *Ephemeral {
	return &Ephemeral{
		userStorage: make(map[uuid.UUID]types.User),
	}
}

func (e Ephemeral) Connect() error {
	return nil
}

func (e Ephemeral) AddUser(user types.User) error {
	e.userStorage[user.ID] = user
	return nil
}

func (e Ephemeral) DeleteUser(user types.User) error {
	//TODO implement me
	panic("implement me")
}

func (e Ephemeral) GetAllUsers() ([]types.User, error) {
	users := make([]types.User, len(e.userStorage))
	i := 0
	for _, user := range e.userStorage {
		users[i] = user
		i++
	}
	return users, nil
}
