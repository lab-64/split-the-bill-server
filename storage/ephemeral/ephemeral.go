package ephemeral

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
	"sync"
)

// TODO: Can change types_test to pointer of types_test -> updating types_test will be displayed in all related structs but it will still not update changes in PendingGroupInvitations
type Ephemeral struct {
	Lock      sync.Mutex
	Users     map[uuid.UUID]model.UserModel
	NameIndex map[string]uuid.UUID
	Passwords map[uuid.UUID][]byte
	Cookies   map[uuid.UUID][]model.AuthCookieModel
	Groups    map[uuid.UUID]*model.GroupModel
	Bills     map[uuid.UUID]*model.BillModel
}

func NewEphemeral() (*Ephemeral, error) {
	return &Ephemeral{
		Users:     make(map[uuid.UUID]model.UserModel),
		NameIndex: make(map[string]uuid.UUID),
		Passwords: make(map[uuid.UUID][]byte),
		Cookies:   make(map[uuid.UUID][]model.AuthCookieModel),
		Groups:    make(map[uuid.UUID]*model.GroupModel),
		Bills:     make(map[uuid.UUID]*model.BillModel),
	}, nil
}
