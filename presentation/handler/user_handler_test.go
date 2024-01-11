package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"split-the-bill-server/authentication"
	"split-the-bill-server/domain/service/mocks"
	"split-the-bill-server/domain/service/service_inf"
	"split-the-bill-server/presentation/dto"
	. "split-the-bill-server/storage/database/test_util"
	"testing"
)

var (
	userService service_inf.IUserService
	userHandler UserHandler
	app         *fiber.App
)

func TestMain(m *testing.M) {
	// setup user handler
	userService = mocks.NewUserServiceMock()
	// password validator
	passwordValidator, err := authentication.NewPasswordValidator()
	if err != nil {
		panic("Error while setting up the password validator: " + err.Error())
	}
	userHandler = *NewUserHandler(&userService, passwordValidator)

	// TODO: setup authentication handler

	// setup fiber
	app = fiber.New()
	// TODO: use SetupRoutes
	app.Get("/user/:id", userHandler.GetByID)
	app.Post("/api/user/register", userHandler.Register)

	// Run tests
	exitCode := m.Run()

	// Exit with the same code as the test run
	os.Exit(exitCode)
}

func TestGetByIDSuccess(t *testing.T) {

	// mock method
	mocks.MockUserGetByID = func(id uuid.UUID) (dto.UserDetailedOutputDTO, error) {
		return dto.ToUserDetailedDTO(&User), nil
	}

	// setup request
	req, _ := http.NewRequest(http.MethodGet, "/user/"+User.ID.String(), nil)
	resp, err := app.Test(req, -1)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		panic(err)
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
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
	assert.EqualValues(t, User.ID, returnedUser.ID)
	assert.EqualValues(t, User.Email, returnedUser.Email)
}

func TestRegisterSuccess(t *testing.T) {

	// mock method
	mocks.MockUserCreate = func(user dto.UserInputDTO) (dto.UserCoreOutputDTO, error) {
		return dto.ToUserCoreDTO(&User), nil
	}

	// setup request
	reqBody := dto.UserInputDTO{
		Email:    User.Email,
		Password: Password,
	}
	jsonBody, err := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, "/api/user/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req, -1)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		panic(err)
	}

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
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
	assert.EqualValues(t, User.ID, returnedUser.ID)
	assert.EqualValues(t, User.Email, returnedUser.Email)
}
