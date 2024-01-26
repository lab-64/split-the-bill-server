package handler

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"split-the-bill-server/domain/model"
	"split-the-bill-server/domain/service/mocks"
	"split-the-bill-server/presentation/dto"
	"testing"
)

// Testdata
var (
	TestUser     = model.UserModel{ID: uuid.New(), Email: "test@mail.com"}
	TestPassword = "test1337"
)

func TestGetByIDSuccess(t *testing.T) {

	// mock method
	mocks.MockUserGetByID = func(id uuid.UUID) (dto.UserDetailedOutputDTO, error) {
		return dto.ConvertToUserDetailedDTO(&TestUser), nil
	}

	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/user/"+TestUser.ID.String(), nil)
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
	var response dto.GeneralResponseDTO
	err = json.Unmarshal(body, &response)

	// convert returned data to dto.User
	var returnedUser dto.UserDetailedOutputDTO
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

func TestRegisterSuccess(t *testing.T) {

	// mock method
	mocks.MockUserCreate = func(user dto.UserInputDTO) (dto.UserCoreOutputDTO, error) {
		return dto.ConvertToUserCoreDTO(&TestUser), nil
	}

	// setup request
	reqBody := dto.UserInputDTO{
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
	var response dto.GeneralResponseDTO
	err = json.Unmarshal(body, &response)

	// convert returned data to dto.User
	var returnedUser dto.UserCoreOutputDTO
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
