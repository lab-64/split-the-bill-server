package ephemeral

import (
	"github.com/google/uuid"
	"split-the-bill-server/domain/model"
)

// TODO: Can change types_test to pointer of types_test -> updating types_test will be displayed in all related structs but it will still not update changes in PendingGroupInvitations
type Ephemeral struct {
	Locker    *Locker
	Users     map[uuid.UUID]model.User
	NameIndex map[string]uuid.UUID
	Passwords map[uuid.UUID][]byte
	Cookies   map[uuid.UUID][]model.AuthCookie
	Groups    map[uuid.UUID]*model.Group
	Bills     map[uuid.UUID]*model.Bill
}

type Resource uint

const (
	RUsers     Resource = 0x1 << iota
	RNameIndex Resource = 0x1 << iota
	RPasswords Resource = 0x1 << iota
	RCookies   Resource = 0x1 << iota
	RGroups    Resource = 0x1 << iota
	RBills     Resource = 0x1 << iota

	// KEEP THIS AS LAST LINE
	NumResources uint = iota
)
const _ = 1 / (64 / NumResources) // compile time check that there are no more than 64 resources

func NewEphemeral() (*Ephemeral, error) {
	return &Ephemeral{
		Locker:    NewLocker(),
		Users:     make(map[uuid.UUID]model.User),
		NameIndex: make(map[string]uuid.UUID),
		Passwords: make(map[uuid.UUID][]byte),
		Cookies:   make(map[uuid.UUID][]model.AuthCookie),
		Groups:    make(map[uuid.UUID]*model.Group),
		Bills:     make(map[uuid.UUID]*model.Bill),
	}, nil
}
