package integration_tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"split-the-bill-server/domain"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/handler"
	"split-the-bill-server/presentation/middleware"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database/entity"
	"testing"
)

type UserResponseDTO struct {
	Message string                `json:"message"`
	Data    dto.UserCoreOutputDTO `json:"data"`
}

// performUserRequest performs a http request for any user endpoint.
// The method, the route including the parameter, the input data and the mail of the user to be logged in must be provided.
// If the request should be performed without authentication token, just provide an empty string as mail.
// The function returns the response body as UserResponseDTO, the http response and an error.
func performUserRequest(httpMethod string, route string, inputUserData interface{}, mailForLogin string) (UserResponseDTO, *http.Response, error) {

	// Get Authentication Token
	cookieToken, setupErr := login(mailForLogin, Password)
	if setupErr != nil {
		return UserResponseDTO{}, nil, setupErr
	}
	cookie := &http.Cookie{Name: "session_cookie", Value: cookieToken}

	inputJSON, _ := json.Marshal(inputUserData)
	// Create http request
	req := httptest.NewRequest(httpMethod, route, bytes.NewReader(inputJSON))
	req.Header.Set("Content-Type", "application/json")
	// add cookie to request
	req.AddCookie(cookie)

	// Perform request
	resp, err := app.Test(req, -1)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return UserResponseDTO{}, resp, err
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserResponseDTO{}, resp, errors.New(fmt.Sprintf("error during reading response body - resp.Body: %s - error: %s", resp.Body, err.Error()))
	}
	// Parse response body to UserResponseDTO
	var response UserResponseDTO
	if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
		return UserResponseDTO{}, resp, errors.New(fmt.Sprintf("error during parsing response - body: %s - error: %s", string(body), err.Error()))
	}
	return response, resp, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Start of CreateUser test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestCreateUser(t *testing.T) {

	route := "/api/user"
	tests := []struct {
		inputData          dto.UserInputDTO      // input data for the request
		description        string                // description of the testcase case
		expectedCode       int                   // expected HTTP status code
		expectedMessage    string                // expected message in response body
		expectReturn       bool                  // expected return value
		expectReturnedData dto.UserCoreOutputDTO // expected return
	}{
		{
			description: "Test successful user creation",
			inputData: dto.UserInputDTO{
				Email:    "test3@mail.com",
				Password: "alek1337",
			},
			expectedCode:    201,
			expectedMessage: handler.SuccessMsgUserCreate,
			expectReturn:    true,
			expectReturnedData: dto.UserCoreOutputDTO{
				Email: "test3@mail.com",
			},
		},
		{
			description: "Test user already exists",
			inputData: dto.UserInputDTO{
				Email:    "test3@mail.com",
				Password: "alek1337",
			},
			expectedCode:    500,
			expectedMessage: fmt.Sprintf(handler.ErrMsgUserCreate, storage.InvalidUserInputError),
			expectReturn:    false,
		},
	}

	// Iterate through testcase single testcase cases
	for _, testcase := range tests {
		responseData, httpResponse, err := performUserRequest(http.MethodPost, route, testcase.inputData, "")
		if err != nil {
			t.Fatalf("Error during setup while performing request: %s", err.Error())
		}

		// Verify, if test case is successfully passed
		assert.Equalf(t, testcase.expectedCode, httpResponse.StatusCode, testcase.description)
		assert.Equalf(t, testcase.expectedMessage, responseData.Message, testcase.description)
		if testcase.expectReturn {
			assert.Equalf(t, testcase.expectReturnedData.Email, responseData.Data.Email, testcase.description)
			// get the stored user from the storage to check if user is correctly stored
			storedUser, setupErr := getStoredUserEntity(responseData.Data.ID)
			assert.Nilf(t, setupErr, "Error during setup while getting stored user entity in test: %s", testcase.description)
			// validate updated user in storage
			assert.Equalf(t, testcase.inputData.Email, storedUser.Email, testcase.description)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Start of GetUser test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestGetUser(t *testing.T) {

	route := "/api/user/"
	tests := []struct {
		description        string                // description of the testcase case
		parameter          string                // parameter of the testcase
		loggedInUser       entity.User           // user that is logged in
		expectedCode       int                   // expected HTTP status code
		expectedMessage    string                // expected message in response body
		expectReturn       bool                  // expected return value
		expectReturnedData dto.UserCoreOutputDTO // expected return
	}{
		{
			description:     "Test successful user query",
			parameter:       User1.ID.String(),
			loggedInUser:    User1,
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgUserFound,
			expectReturn:    true,
			expectReturnedData: dto.UserCoreOutputDTO{
				ID:    User1.ID,
				Email: User1.Email,
			},
		},
		// TODO: maybe add a test case for missing cookie
		{
			description:     "Test auth cookie is invalid",
			parameter:       User1.ID.String(),
			loggedInUser:    entity.User{},
			expectedCode:    401,
			expectedMessage: middleware.ErrMsgInvalidCookie,
			expectReturn:    false,
		},
		// TODO: Can different persons query the user?
		/*		{
				description:     "Test user is unauthorized",
				parameter:       User1.ID.String(),
				loggedInUser:    User2,
				expectedCode:    401,
				expectedMessage: fmt.Sprintf(middleware.ErrMsgAuthentication, storage.NoSuchCookieError),
				expectReturn:    false,
			},*/
	}

	// Iterate through testcase single testcase cases
	for _, testcase := range tests {
		responseData, httpResponse, err := performUserRequest(http.MethodGet, route+testcase.parameter, nil, testcase.loggedInUser.Email)
		if err != nil {
			t.Fatalf("Error during setup while performing request: %s", err.Error())
		}

		// Verify, if test case is successfully passed
		assert.Equalf(t, testcase.expectedCode, httpResponse.StatusCode, testcase.description) // check status code
		assert.Equalf(t, testcase.expectedMessage, responseData.Message, testcase.description) // check message
		if testcase.expectReturn {
			assert.Equalf(t, testcase.expectReturnedData.ID, responseData.Data.ID, testcase.description)       // check returned id
			assert.Equalf(t, testcase.expectReturnedData.Email, responseData.Data.Email, testcase.description) // check returned mail
		}

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Start of UpdateUser test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestUpdateUser(t *testing.T) {

	route := "/api/user/"
	tests := []struct {
		description        string
		loggedInUser       entity.User
		parameter          uuid.UUID
		inputUser          dto.UserUpdateDTO
		expectedCode       int
		expectedMessage    string
		expectReturn       bool
		expectReturnedData dto.UserCoreOutputDTO
	}{
		{
			description:  "Test successful user update",
			loggedInUser: User1,
			parameter:    User1.ID,
			inputUser: dto.UserUpdateDTO{
				Email:    "new-mail@mail.com",
				Username: "Franz",
			},
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgUserUpdate,
			expectReturn:    true,
			expectReturnedData: dto.UserCoreOutputDTO{
				ID:       User1.ID,
				Email:    "new-mail@mail.com",
				Username: "Franz",
			},
		},
		{
			description:     "Test unsuccessful behavior: user is unauthorized to update foreign user",
			loggedInUser:    User2,
			parameter:       User1.ID,
			inputUser:       dto.UserUpdateDTO{},
			expectedCode:    401,
			expectedMessage: fmt.Sprintf(handler.ErrMsgUserUpdate, domain.ErrNotAuthorized),
			expectReturn:    false,
		},
		{
			description:     "Test unsuccessful behavior: user is not logged in",
			loggedInUser:    entity.User{},
			parameter:       User1.ID,
			inputUser:       dto.UserUpdateDTO{},
			expectedCode:    401,
			expectedMessage: middleware.ErrMsgInvalidCookie,
			expectReturn:    false,
		},
	}
	for _, testcase := range tests {
		responseData, httpResponse, err := performUserRequest(http.MethodPut, route+testcase.parameter.String(), testcase.inputUser, testcase.loggedInUser.Email)
		if err != nil {
			t.Fatalf("Error during setup while performing request: %s", err.Error())
		}

		assert.Equalf(t, testcase.expectedCode, httpResponse.StatusCode, testcase.description)
		assert.Equalf(t, testcase.expectedMessage, responseData.Message, testcase.description)
		if testcase.expectReturn {
			// validate response
			assert.Equalf(t, testcase.parameter, responseData.Data.ID, testcase.description) // parameter contains the id of the issuer
			assert.Equalf(t, testcase.inputUser.Email, responseData.Data.Email, testcase.description)
			assert.Equalf(t, testcase.inputUser.Username, responseData.Data.Username, testcase.description)
			// get the stored user from the storage to check if user is correctly stored
			storedUser, setupErr := getStoredUserEntity(testcase.parameter)
			assert.Nilf(t, setupErr, "Error during setup while getting stored user entity in test: %s", testcase.description)
			// validate updated user in storage
			assert.Equalf(t, testcase.parameter, storedUser.ID, testcase.description)
			assert.Equalf(t, testcase.inputUser.Email, storedUser.Email, testcase.description)
			assert.Equalf(t, testcase.inputUser.Username, storedUser.Username, testcase.description)
		}
	}
}
