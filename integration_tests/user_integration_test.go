package integration_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http/httptest"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/handler"
	"split-the-bill-server/storage"
	"testing"
)

func TestCreateUser(t *testing.T) {
	refreshDB()

	tests := []struct {
		inputJSON       string
		description     string // description of the testcase case
		route           string // route path to testcase
		expectedCode    int    // expected HTTP status code
		expectedMessage string // expected message in response body
		expectReturn    bool   // expected return value
	}{
		// test successful user creation
		{
			description:     "Test successful user creation",
			inputJSON:       `{"email": "test3@mail.com", "password": "alek1337"}`,
			route:           "/api/user/register",
			expectedCode:    201,
			expectedMessage: handler.SuccessMsgUserCreate,
			expectReturn:    true,
		},
		// test user already exist
		{
			description:     "Test user already exists",
			inputJSON:       `{"email": "test3@mail.com", "password": "alek1337"}`,
			route:           "/api/user/register",
			expectedCode:    500,
			expectedMessage: fmt.Sprintf(handler.ErrMsgUserCreate, storage.InvalidUserInputError),
			expectReturn:    false,
		},
	}

	// Iterate through testcase single testcase cases
	for _, testcase := range tests {
		// Create a new http request with the route from the testcase case
		req := httptest.NewRequest("POST", testcase.route, bytes.NewBufferString(testcase.inputJSON))
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, err := app.Test(req, -1)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			log.Println("test err")
			panic(err)
		}

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("read err")
			panic(err)
		}
		// Parse response body to GeneralResponseDTO
		var response dto.GeneralResponseDTO
		if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
			log.Println("unmarshal err")
			panic(err)
		}
		log.Println(response.Message)

		// Verify, if test case is successfully passed
		assert.Equalf(t, testcase.expectedCode, resp.StatusCode, testcase.description)      // check status code
		assert.Equalf(t, testcase.expectedMessage, response.Message, testcase.description)  // check message
		assert.Equalf(t, testcase.expectReturn, response.Data != nil, testcase.description) // check returned data
	}
}
