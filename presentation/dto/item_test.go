package dto

import (
	"github.com/google/uuid"
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
				Name:   "Test",
				Price:  10.0,
				BillID: uuid.New(),
			},
			expectedErr: nil,
		},
		{
			name: "Name is missing",
			input: ItemInput{
				Price:  10.0,
				BillID: uuid.New(),
			},
			expectedErr: ErrItemNameRequired,
		},
		{
			name: "Price is missing",
			input: ItemInput{
				Name:   "Test",
				BillID: uuid.New(),
			},
			expectedErr: ErrItemPriceRequired,
		},
		{
			name: "BillID is missing",
			input: ItemInput{
				Name:  "Test",
				Price: 10.0,
			},
			expectedErr: ErrItemBillIDRequired,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			err := testcase.input.ValidateInputs()
			assert.Equalf(t, testcase.expectedErr, err, "Wrong error")
		})
	}
}
