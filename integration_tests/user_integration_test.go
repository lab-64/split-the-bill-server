package integration_tests

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
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/handler"
	"split-the-bill-server/presentation/middleware"
	"split-the-bill-server/storage"
	"split-the-bill-server/storage/database/entity"
	"testing"
)

type UserResponseDTO struct {
	Message string             `json:"message"`
	Data    dto.UserCoreOutput `json:"data"`
}

// performUserRequest performs a http request for any user endpoint.
// The method, the route including the parameter, the input data and the session cookie has to be provided.
// The function returns the response body as UserResponseDTO, the http response and an error.
func performUserRequest(httpMethod string, route string, inputUserData interface{}, cookie *http.Cookie) (UserResponseDTO, *http.Response, error) {

	inputJSON, _ := json.Marshal(inputUserData)
	// Create http request
	req := httptest.NewRequest(httpMethod, route, bytes.NewReader(inputJSON))
	req.Header.Set("Content-Type", "application/json")
	// add cookie to request if set
	if cookie != nil {
		req.AddCookie(cookie)
	}

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
		return UserResponseDTO{}, resp, fmt.Errorf("error during reading response body - resp.Body: %s - error: %s", resp.Body, err.Error())
	}
	// Parse response body to UserResponseDTO
	var response UserResponseDTO
	if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
		return UserResponseDTO{}, resp, fmt.Errorf("error during parsing response - body: %s - error: %s", string(body), err.Error())
	}
	return response, resp, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Start of CreateUser test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestCreateUser(t *testing.T) {

	route := "/api/user"
	tests := []struct {
		description        string             // description of the testcase case
		inputData          dto.UserInput      // input data for the request
		requestCookie      *http.Cookie       // cookie for the request
		expectedCode       int                // expected HTTP status code
		expectedMessage    string             // expected message in response body
		expectReturn       bool               // expected return value
		expectReturnedData dto.UserCoreOutput // expected return
	}{
		{
			description: "Test successful user creation",
			inputData: dto.UserInput{
				Email:    "test3@mail.com",
				Password: "alek1337",
			},
			requestCookie:   nil,
			expectedCode:    201,
			expectedMessage: handler.SuccessMsgUserCreate,
			expectReturn:    true,
			expectReturnedData: dto.UserCoreOutput{
				Email: "test3@mail.com",
			},
		},
		{
			description: "Test user already exists",
			inputData: dto.UserInput{
				Email:    "test3@mail.com",
				Password: "alek1337",
			},
			requestCookie:   nil,
			expectedCode:    500,
			expectedMessage: fmt.Sprintf(handler.ErrMsgUserCreate, storage.InvalidUserInputError),
			expectReturn:    false,
		},
	}

	// Iterate through testcase single testcase cases
	for _, testcase := range tests {
		responseData, httpResponse, err := performUserRequest(http.MethodPost, route, testcase.inputData, testcase.requestCookie)
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
		description        string             // description of the testcase case
		parameter          string             // parameter of the testcase
		requestCookie      *http.Cookie       // cookie for the request
		expectedCode       int                // expected HTTP status code
		expectedMessage    string             // expected message in response body
		expectReturn       bool               // expected return value
		expectReturnedData dto.UserCoreOutput // expected return
	}{
		{
			description:     "Test successful user query",
			parameter:       User1.ID.String(),
			requestCookie:   &http.Cookie{Name: sessionCookie, Value: CookieUser1.ID.String()},
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgUserFound,
			expectReturn:    true,
			expectReturnedData: dto.UserCoreOutput{
				ID:    User1.ID,
				Email: User1.Email,
			},
		},
		{
			description:     "Test auth cookie is missing",
			parameter:       User1.ID.String(),
			requestCookie:   nil,
			expectedCode:    401,
			expectedMessage: middleware.ErrMsgNoCookie,
			expectReturn:    false,
		},
	}

	// Iterate through testcase single testcase cases
	for _, testcase := range tests {
		responseData, httpResponse, err := performUserRequest(http.MethodGet, route+testcase.parameter, nil, testcase.requestCookie)
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
		requester          entity.User
		parameter          uuid.UUID
		inputUser          dto.UserUpdate
		requestCookie      *http.Cookie
		expectedCode       int
		expectedMessage    string
		expectReturn       bool
		expectReturnedData dto.UserCoreOutput
	}{
		{
			description: "Test successful user update",
			requester:   User1,
			parameter:   User1.ID,
			inputUser: dto.UserUpdate{
				Username: "Franz",
			},
			requestCookie:   &http.Cookie{Name: sessionCookie, Value: CookieUser1.ID.String()},
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgUserUpdate,
			expectReturn:    true,
			expectReturnedData: dto.UserCoreOutput{
				ID:       User1.ID,
				Email:    "new-mail@mail.com",
				Username: "Franz",
			},
		},
		{
			description:     "Test unsuccessful behavior: user is unauthorized to update foreign user",
			requester:       User2,
			parameter:       User1.ID,
			inputUser:       dto.UserUpdate{},
			requestCookie:   &http.Cookie{Name: sessionCookie, Value: CookieUser2.ID.String()},
			expectedCode:    401,
			expectedMessage: fmt.Sprintf(handler.ErrMsgUserUpdate, domain.ErrNotAuthorized),
			expectReturn:    false,
		},
		{
			description:     "Test unsuccessful behavior: user is not logged in",
			requester:       entity.User{},
			parameter:       User1.ID,
			inputUser:       dto.UserUpdate{},
			requestCookie:   nil,
			expectedCode:    401,
			expectedMessage: middleware.ErrMsgNoCookie,
			expectReturn:    false,
		},
	}
	for _, testcase := range tests {
		responseData, httpResponse, err := performUserRequest(http.MethodPut, route+testcase.parameter.String(), testcase.inputUser, testcase.requestCookie)
		if err != nil {
			t.Fatalf("Error during setup while performing request: %s", err.Error())
		}

		assert.Equalf(t, testcase.expectedCode, httpResponse.StatusCode, testcase.description)
		assert.Equalf(t, testcase.expectedMessage, responseData.Message, testcase.description)
		if testcase.expectReturn {
			// validate response
			assert.Equalf(t, testcase.parameter, responseData.Data.ID, testcase.description) // parameter contains the id of the issuer
			assert.Equalf(t, testcase.requester.Email, responseData.Data.Email, testcase.description)
			assert.Equalf(t, testcase.inputUser.Username, responseData.Data.Username, testcase.description)
			// get the stored user from the storage to check if user is correctly stored
			storedUser, setupErr := getStoredUserEntity(testcase.parameter)
			assert.Nilf(t, setupErr, "Error during setup while getting stored user entity in test: %s", testcase.description)
			// validate updated user in storage
			assert.Equal(t, testcase.parameter, storedUser.ID)
			assert.Equal(t, testcase.requester.Email, storedUser.Email) // email should not be changed
			assert.Equal(t, testcase.inputUser.Username, storedUser.Username)
		}
	}
}
