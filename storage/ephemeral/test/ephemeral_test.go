package test

import (
	"split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/ephemeral/eph_storages"
	"testing"
)

func TestEphemeral(t *testing.T) {
	e, _ := ephemeral.NewEphemeral()
	userStorage := eph_storages.NewUserStorage(e)
	UserStorageTest(e, userStorage, t)
}

func TestEphemeralEdgeCases(t *testing.T) {
	e, _ := ephemeral.NewEphemeral()
	userStorage := eph_storages.NewUserStorage(e)
	UserStorageEdgeCaseTest(e, userStorage, t)
}
