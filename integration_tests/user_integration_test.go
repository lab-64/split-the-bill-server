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
	Message string                `json:"message"`
	Data    dto.UserCoreOutputDTO `json:"data"`
}

func TestCreateUser(t *testing.T) {

	route := "/api/user"
	tests := []struct {
		inputJSON          string
		description        string                // description of the testcase case
		expectedCode       int                   // expected HTTP status code
		expectedMessage    string                // expected message in response body
		expectReturn       bool                  // expected return value
		expectReturnedData dto.UserCoreOutputDTO // expected return
	}{
		{
			description:     "Test successful user creation",
			inputJSON:       `{"email": "test3@mail.com", "password": "alek1337"}`,
			expectedCode:    201,
			expectedMessage: handler.SuccessMsgUserCreate,
			expectReturn:    true,
			expectReturnedData: dto.UserCoreOutputDTO{
				ID:    uuid.New(),
				Email: "test3@mail.com",
			},
		},
		{
			description:     "Test user already exists",
			inputJSON:       `{"email": "test3@mail.com", "password": "alek1337"}`,
			expectedCode:    500,
			expectedMessage: fmt.Sprintf(handler.ErrMsgUserCreate, storage.InvalidUserInputError),
			expectReturn:    false,
		},
	}

	// Iterate through testcase single testcase cases
	for _, testcase := range tests {
		// Create http request
		req := httptest.NewRequest("POST", route, bytes.NewBufferString(testcase.inputJSON))
		req.Header.Set("Content-Type", "application/json")

		// Perform request
		resp, err := app.Test(req, -1)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			t.Fatal("Error in test setup during performing a request - ", err)
		}

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("Error in test setup during reading response body - ", err)
		}
		// Parse response body to GeneralResponseDTO
		var response dto.GeneralResponseDTO
		if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
			t.Fatal("Error in test setup during parsing response body - ", err)
		}

		// Verify, if test case is successfully passed
		assert.Equalf(t, testcase.expectedCode, resp.StatusCode, testcase.description)      // check status code
		assert.Equalf(t, testcase.expectedMessage, response.Message, testcase.description)  // check message
		assert.Equalf(t, testcase.expectReturn, response.Data != nil, testcase.description) // check returned data
		if testcase.expectReturn {
			returnData := response.Data.(map[string]interface{})
			assert.Equal(t, testcase.expectReturnedData.Email, returnData["email"]) // check returned id
		}
	}
}

func TestGetUser(t *testing.T) {

	route := "/api/user/"
	tests := []struct {
		description        string                // description of the testcase case
		parameter          string                // parameter of the testcase
		cookie             *http.Cookie          // cookie of the testcase
		expectedCode       int                   // expected HTTP status code
		expectedMessage    string                // expected message in response body
		expectReturn       bool                  // expected return value
		expectReturnedData dto.UserCoreOutputDTO // expected return
	}{
		{
			description:     "Test successful user query",
			parameter:       User1.ID.String(),
			cookie:          &http.Cookie{Name: "session_cookie", Value: Cookie1.ID.String()},
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgUserFound,
			expectReturn:    true,
			expectReturnedData: dto.UserCoreOutputDTO{
				ID:    User1.ID,
				Email: User1.Email,
			},
		},
		{
			description:     "Test auth cookie is missing",
			parameter:       User1.ID.String(),
			cookie:          nil,
			expectedCode:    401,
			expectedMessage: middleware.ErrMsgNoCookie,
			expectReturn:    false,
		},
		{
			description:     "Test user is unauthorized",
			parameter:       User1.ID.String(),
			cookie:          &http.Cookie{Name: "session_cookie", Value: uuid.NewString()},
			expectedCode:    401,
			expectedMessage: fmt.Sprintf(middleware.ErrMsgAuthentication, storage.NoSuchCookieError),
			expectReturn:    false,
		},
	}

	// Iterate through testcase single testcase cases
	for _, testcase := range tests {
		// Create http request
		req := httptest.NewRequest(http.MethodGet, route+testcase.parameter, nil)
		if testcase.cookie != nil {
			req.AddCookie(testcase.cookie)
		}
		// Perform request
		resp, err := app.Test(req, -1)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			t.Fatal("Error in test setup during performing a request - ", err)
		}

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("Error in test setup during reading response body - ", err)
		}
		// Parse response body to GeneralResponseDTO
		var response dto.GeneralResponseDTO
		if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
			t.Fatal("Error in test setup during parsing response body - ", err)
		}

		// Verify, if test case is successfully passed
		assert.Equal(t, testcase.expectedCode, resp.StatusCode)                             // check status code
		assert.Equal(t, testcase.expectedMessage, response.Message)                         // check message
		assert.Equalf(t, testcase.expectReturn, response.Data != nil, testcase.description) // check returned data
		if testcase.expectReturn {

			// convert returned data to dto.User
			returnData := response.Data.(map[string]interface{})
			var returnedUser dto.UserCoreOutputDTO
			returnedUser.ID = uuid.MustParse(returnData["id"].(string))
			returnedUser.Email = returnData["email"].(string)
			assert.Equal(t, testcase.expectReturnedData.ID, returnedUser.ID)       // check returned id
			assert.Equal(t, testcase.expectReturnedData.Email, returnedUser.Email) // check returned mail
		}

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Start of UpdateUser test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func performUpdateUserRequest(parameter uuid.UUID, testInputUser dto.UserUpdateDTO, username string) (UserResponseDTO, *http.Response, error) {
	route := "/api/user/"

	// Get Authentication Token
	cookieToken, setupErr := login(username, Password)
	if setupErr != nil {
		return UserResponseDTO{}, nil, setupErr
	}

	inputJSON, _ := json.Marshal(testInputUser)
	cookie := &http.Cookie{Name: "session_cookie", Value: cookieToken}

	// Create http request
	req := httptest.NewRequest("PUT", route+parameter.String(), bytes.NewReader(inputJSON))
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
		return UserResponseDTO{}, resp, err
	}
	// Parse response body to GeneralResponseDTO
	var response UserResponseDTO
	if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
		return UserResponseDTO{}, resp, err
	}
	return response, resp, nil
}

func TestUpdateUser(t *testing.T) {

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
		responseData, httpResponse, err := performUpdateUserRequest(testcase.parameter, testcase.inputUser, testcase.loggedInUser.Email)
		if err != nil {
			t.Fatalf("Error during setup while performing request: %s", err.Error())
		}

		assert.Nil(t, err)
		assert.Equal(t, testcase.expectedCode, httpResponse.StatusCode)
		assert.Equal(t, testcase.expectedMessage, responseData.Message)
		if testcase.expectReturn {
			// validate response
			assert.Equal(t, testcase.parameter, responseData.Data.ID) // parameter contains the id of the issuer
			assert.Equal(t, testcase.inputUser.Username, responseData.Data.Username)
			// get the stored user from the storage for comparison
			storedUser, _ := getStoredUserEntity(testcase.parameter)
			// validate updated user in storage
			assert.Equal(t, testcase.parameter, storedUser.ID)
			assert.Equal(t, testcase.loggedInUser.Email, storedUser.Email) // email should not be changed
			assert.Equal(t, testcase.inputUser.Username, storedUser.Username)
		}
	}
}
