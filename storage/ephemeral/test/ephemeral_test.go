package test

import (
	"split-the-bill-server/storage/ephemeral"
	"testing"
)

func TestEphemeral(t *testing.T) {
	e, _ := ephemeral.NewEphemeral()
	userStorage := ephemeral.NewUserStorage(e)
	UserStorageTest(e, userStorage, t)
}

func TestEphemeralEdgeCases(t *testing.T) {
	e, _ := ephemeral.NewEphemeral()
	userStorage := ephemeral.NewUserStorage(e)
	UserStorageEdgeCaseTest(e, userStorage, t)
}
