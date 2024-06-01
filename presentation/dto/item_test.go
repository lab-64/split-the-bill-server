package dto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItemInput_ValidateInputs(t *testing.T) {
	tests := []struct {
		name        string
		input       ItemInput
		expectedErr error
	}{
		{
			name: "Success",
			input: ItemInput{
				Name:  "Test",
				Price: 10.0,
			},
			expectedErr: nil,
		},
		{
			name: "Name is missing",
			input: ItemInput{
				Price: 10.0,
			},
			expectedErr: ErrItemNameRequired,
		},
		{
			name: "Price is missing",
			input: ItemInput{
				Name: "Test",
			},
			expectedErr: ErrItemPriceRequired,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			err := testcase.input.ValidateInputs()
			assert.Equalf(t, testcase.expectedErr, err, "Wrong error")
		})
	}
}
