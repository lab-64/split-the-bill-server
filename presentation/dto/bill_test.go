package dto

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBillInput_ValidateInputs(t *testing.T) {
	tests := []struct {
		name        string
		input       BillInput
		expectedErr error
	}{
		{
			name: "Success",
			input: BillInput{
				OwnerID: uuid.New(),
				Name:    "Test",
				Date:    time.Now(),
				GroupID: uuid.New(),
			},
			expectedErr: nil,
		},
		{
			name: "OwnerID is missing",
			input: BillInput{
				Name:    "Test",
				Date:    time.Now(),
				GroupID: uuid.New(),
			},
			expectedErr: ErrBillOwnerIDRequired,
		},
		{
			name: "Name is missing",
			input: BillInput{
				OwnerID: uuid.New(),
				Date:    time.Now(),
				GroupID: uuid.New(),
			},
			expectedErr: ErrBillNameRequired,
		},
		{
			name: "Date is missing",
			input: BillInput{
				OwnerID: uuid.New(),
				Name:    "Test",
				GroupID: uuid.New(),
			},
			expectedErr: ErrBillDateRequired,
		},
		{
			name: "GroupID is missing",
			input: BillInput{
				OwnerID: uuid.New(),
				Name:    "Test",
				Date:    time.Now(),
			},
			expectedErr: ErrBillGroupIDRequired,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			err := testcase.input.ValidateInputs()
			assert.Equalf(t, testcase.expectedErr, err, "Wrong error")
		})
	}
}
