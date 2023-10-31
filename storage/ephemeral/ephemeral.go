package ephemeral

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"sync"
)

// TODO: Can change types_test to pointer of types_test -> updating types_test will be displayed in all related structs but it will still not update changes in PendingGroupInvitations
type Ephemeral struct {
	Lock      sync.Mutex
	Users     map[uuid.UUID]model.User
	NameIndex map[string]uuid.UUID
	Passwords map[uuid.UUID][]byte
	Cookies   map[uuid.UUID][]model.AuthenticationCookie
	Groups    map[uuid.UUID]*model.Group
	Bills     map[uuid.UUID]*model.Bill
}

func NewEphemeral() (*Ephemeral, error) {
	return &Ephemeral{
		Users:     make(map[uuid.UUID]model.User),
		NameIndex: make(map[string]uuid.UUID),
		Passwords: make(map[uuid.UUID][]byte),
		Cookies:   make(map[uuid.UUID][]model.AuthenticationCookie),
		Groups:    make(map[uuid.UUID]*model.Group),
		Bills:     make(map[uuid.UUID]*model.Bill),
	}, nil
}

func (e *Ephemeral) Connect() error {
	return nil
}
