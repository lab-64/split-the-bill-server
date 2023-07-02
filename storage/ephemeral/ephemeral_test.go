package ephemeral_test

import (
	"split-the-bill-server/storage/ephemeral"
	"split-the-bill-server/storage/test"
	"testing"
)

func TestEphemeral(t *testing.T) {
	uut := ephemeral.NewEphemeral()
	test.UserStorageTest(uut, t)
}

func TestEphemeralEdgeCases(t *testing.T) {
	uut := ephemeral.NewEphemeral()
	test.UserStorageEdgeCaseTest(uut, t)
}
