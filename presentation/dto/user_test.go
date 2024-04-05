package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserInput_ValidateInputs(t *testing.T) {
	tests := []struct {
		name        string
		input       UserInput
		expectedErr error
	}{
		{
			name: "Success",
			input: UserInput{
				Email:    "test@mail.com",
				Password: "password",
			},
			expectedErr: nil,
		},
		{
			name: "Email is empty",
			input: UserInput{
				Email:    "",
				Password: "password",
			},
			expectedErr: ErrEmailRequired,
		},
		{
			name: "Email is missing",
			input: UserInput{
				Password: "password",
			},
			expectedErr: ErrEmailRequired,
		},
		{
			name: "Password is empty",
			input: UserInput{
				Email:    "test@mail.com",
				Password: "",
			},
			expectedErr: ErrPasswordRequired,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			err := testcase.input.ValidateInputs()
			assert.Equalf(t, testcase.expectedErr, err, "Wrong error")
		})
	}

}

func TestUserUpdate_ValidateInputs(t *testing.T) {
	tests := []struct {
		name        string
		input       UserUpdate
		contentType string
		expectedErr error
	}{
		{
			name: "Success",
			input: UserUpdate{
				Username: "test",
			},
			contentType: "image/jpeg",
			expectedErr: nil,
		},
		{
			name: "Wrong image type",
			input: UserUpdate{
				Username: "test",
			},
			contentType: "text/plain",
			expectedErr: ErrWrongImageType,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			err := testcase.input.ValidateInputs(testcase.contentType)
			assert.Equalf(t, testcase.expectedErr, err, "Wrong error")
		})
	}
}
