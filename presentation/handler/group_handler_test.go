package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"split-the-bill-server/domain/service/mocks"
	"split-the-bill-server/presentation/dto"
	"testing"
)

type GroupResponseDTO struct {
	Message string                  `json:"message"`
	Data    dto.GroupDetailedOutput `json:"data"`
}

func performGroupRequest(httpMethod string, url string, body []byte) (*http.Response, GroupResponseDTO, error) {
	req := httptest.NewRequest(httpMethod, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, GroupResponseDTO{}, fmt.Errorf("error in test setup during performing a request - %w", err)
	}
	// read response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, GroupResponseDTO{}, fmt.Errorf("error in test setup during reading response body - %w", err)
	}
	var response GroupResponseDTO
	err = json.Unmarshal(body, &response)
	return resp, response, err
}

func TestGroupHandler_Create(t *testing.T) {
	route := "/api/group"

	tests := []struct {
		name            string
		requesterID     uuid.UUID
		body            dto.GroupInput
		mock            func()
		expectedStatus  int
		expectedMessage string
	}{
		{
			name:        "Success",
			requesterID: TestUser.ID,
			body: dto.GroupInput{
				OwnerID: TestUser.ID,
				Name:    "Test Group",
			},
			mock: func() {
				mocks.MockGroupCreate = func(requesterID uuid.UUID, groupDTO dto.GroupInput) (dto.GroupDetailedOutput, error) {
					return dto.GroupDetailedOutput{}, nil
				}
			},
			expectedStatus:  http.StatusCreated,
			expectedMessage: SuccessMsgGroupCreate,
		},
		{
			name: "Missing OwnerID",
			body: dto.GroupInput{
				Name: "Test Group",
			},
			mock:            func() {},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: fmt.Sprintf(ErrMsgInputsInvalid, dto.ErrOwnerIDRequired),
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			testcase.mock()
			body, _ := json.Marshal(testcase.body)
			resp, response, err := performGroupRequest(http.MethodPost, route, body)
			assert.Nilf(t, err, "Request failed")
			assert.Equalf(t, testcase.expectedStatus, resp.StatusCode, "Wrong status code")
			assert.Equalf(t, testcase.expectedMessage, response.Message, "Wrong message")
		})
	}
}
