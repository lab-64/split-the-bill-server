package ephemeral

import (
	"github.com/google/uuid"
	"split-the-bill-server/types"
	"sync"
)

// TODO: Can change types_test to pointer of types_test -> updating types_test will be displayed in all related structs but it will still not update changes in PendingGroupInvitations
type Ephemeral struct {
	lock      sync.Mutex
	users     map[uuid.UUID]types.User
	nameIndex map[string]uuid.UUID
	passwords map[uuid.UUID][]byte
	cookies   map[uuid.UUID][]types.AuthenticationCookie
	groups    map[uuid.UUID]*types.Group
	bills     map[uuid.UUID]*types.Bill
}

func NewEphemeral() (*Ephemeral, error) {
	return &Ephemeral{
		users:     make(map[uuid.UUID]types.User),
		nameIndex: make(map[string]uuid.UUID),
		passwords: make(map[uuid.UUID][]byte),
		cookies:   make(map[uuid.UUID][]types.AuthenticationCookie),
		groups:    make(map[uuid.UUID]*types.Group),
		bills:     make(map[uuid.UUID]*types.Bill),
	}, nil
}

func (e *Ephemeral) Connect() error {
	return nil
}
