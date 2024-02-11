package dto

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupInput_ValidateInput(t *testing.T) {
	tests := []struct {
		name        string
		input       GroupInput
		expectedErr error
	}{
		{
			name: "Success",
			input: GroupInput{
				OwnerID: uuid.New(),
				Name:    "Test Group",
			},
			expectedErr: nil,
		},
		{
			name: "OwnerID is missing",
			input: GroupInput{
				Name: "Test Group",
			},
			expectedErr: ErrGroupOwnerIDRequired,
		},
		{
			name: "Name is empty",
			input: GroupInput{
				OwnerID: uuid.New(),
			},
			expectedErr: ErrGroupNameRequired,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			err := testcase.input.ValidateInput()
			assert.Equalf(t, testcase.expectedErr, err, "Wrong error")
		})
	}
}
