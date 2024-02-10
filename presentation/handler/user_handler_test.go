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
	"split-the-bill-server/domain"
	"split-the-bill-server/domain/converter"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/domain/service/mocks"
	"split-the-bill-server/presentation/dto"
	"testing"
)

// Testdata
var (
	TestUser     = model.User{ID: uuid.New(), Email: "test@mail.com", Username: "tester"}
	TestPassword = "test1337"
)

type UserResponseDTO struct {
	Message string             `json:"message"`
	Data    dto.UserCoreOutput `json:"data"`
}

func performRequest(httpMethod string, url string, body []byte) (*http.Response, UserResponseDTO, error) {
	req := httptest.NewRequest(httpMethod, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, UserResponseDTO{}, fmt.Errorf("error in test setup during performing a request - %w", err)
	}
	// read response body
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, UserResponseDTO{}, fmt.Errorf("error in test setup during reading response body - %w", err)
	}
	var response UserResponseDTO
	err = json.Unmarshal(body, &response)
	return resp, response, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Get User test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestGetByIDSuccess(t *testing.T) {

	// mock method
	mocks.MockUserGetByID = func(id uuid.UUID) (dto.UserCoreOutput, error) {
		return converter.ToUserCoreDTO(&TestUser), nil
	}

	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/api/user/"+TestUser.ID.String(), nil)
	resp, err := app.Test(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Fatal("Error in test setup during performing a request - ", err)
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error in test setup during reading response body - ", err)
	}
	var response dto.GeneralResponse
	err = json.Unmarshal(body, &response)

	// convert returned data to dto.User
	var returnedUser dto.UserCoreOutput
	returnData := response.Data.(map[string]interface{})
	returnedUser.ID = uuid.MustParse(returnData["id"].(string))
	returnedUser.Email = returnData["email"].(string)

	// validate test
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, SuccessMsgUserFound, response.Message)
	assert.EqualValues(t, TestUser.ID, returnedUser.ID)
	assert.EqualValues(t, TestUser.Email, returnedUser.Email)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Register User test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestRegisterSuccess(t *testing.T) {

	// mock method
	mocks.MockUserCreate = func(user dto.UserInput) (dto.UserCoreOutput, error) {
		return converter.ToUserCoreDTO(&TestUser), nil
	}

	// setup request
	reqBody := dto.UserInput{
		Email:    TestUser.Email,
		Password: TestPassword,
	}
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		t.Fatal("Error in test setup during performing a request - ", err)
	}
	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error in test setup during reading response body - ", err)
	}
	var response dto.GeneralResponse
	err = json.Unmarshal(body, &response)

	// convert returned data to dto.User
	var returnedUser dto.UserCoreOutput
	returnData := response.Data.(map[string]interface{})
	returnedUser.ID = uuid.MustParse(returnData["id"].(string))
	returnedUser.Email = returnData["email"].(string)

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, SuccessMsgUserCreate, response.Message)
	assert.EqualValues(t, TestUser.ID, returnedUser.ID)
	assert.EqualValues(t, TestUser.Email, returnedUser.Email)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Update User test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestUpdateSuccess(t *testing.T) {

	// mock method
	mocks.MockUserUpdate = func(requesterID uuid.UUID, id uuid.UUID, user dto.UserUpdate) (dto.UserCoreOutput, error) {
		return dto.UserCoreOutput{ID: id, Email: TestUser.Email, Username: user.Username}, nil
	}

	reqBody := dto.UserUpdate{
		Username: "Updated Tester",
	}
	jsonBody, _ := json.Marshal(reqBody)

	resp, response, err := performRequest(http.MethodPut, "/api/user/"+TestUser.ID.String(), jsonBody)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, SuccessMsgUserUpdate, response.Message)
	assert.EqualValues(t, TestUser.ID, response.Data.ID)
	assert.EqualValues(t, TestUser.Email, response.Data.Email)
	assert.EqualValues(t, "Updated Tester", response.Data.Username)
}

func TestUpdateWrongUser(t *testing.T) {

	// mock method
	mocks.MockUserUpdate = func(requesterID uuid.UUID, id uuid.UUID, user dto.UserUpdate) (dto.UserCoreOutput, error) {
		return dto.UserCoreOutput{}, domain.ErrNotAuthorized
	}

	reqBody := dto.UserUpdate{
		Username: TestUser.Username,
	}
	jsonBody, _ := json.Marshal(reqBody)

	resp, response, err := performRequest(http.MethodPut, "/api/user/"+uuid.New().String(), jsonBody)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusUnauthorized, resp.StatusCode)
	assert.Equal(t, fmt.Sprintf(ErrMsgUserUpdate, domain.ErrNotAuthorized.Error()), response.Message)
}
