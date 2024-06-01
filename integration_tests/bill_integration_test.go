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
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/handler"
	"split-the-bill-server/storage/database/entity"
	"testing"
	"time"
)

type BillResponseDTO struct {
	Message string                 `json:"message"`
	Data    dto.BillDetailedOutput `json:"data"`
}

// performBillRequest performs a http request for any user endpoint.
// The method, the route including the parameter, the input data and the session cookie has to be provided.
// The function returns the response body as BillResponseDTO, the http response and an error.
func performBillRequest(httpMethod string, route string, inputUserData interface{}, cookie *http.Cookie) (BillResponseDTO, *http.Response, error) {

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
		return BillResponseDTO{}, resp, err
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return BillResponseDTO{}, resp, fmt.Errorf("error during reading response body - resp.Body: %s - error: %s", resp.Body, err.Error())
	}
	// Parse response body to UserResponseDTO
	var response BillResponseDTO
	if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
		return BillResponseDTO{}, resp, fmt.Errorf("error during parsing response - body: %s - error: %s", string(body), err.Error())
	}
	return response, resp, nil
}

func TestUpdateBill(t *testing.T) {

	// get the bill first
	responseData, _, _ := performBillRequest(http.MethodGet, "/api/bill/"+Bill1.ID.String(), nil, &http.Cookie{Name: sessionCookie, Value: CookieUser1.ID.String()})

	// Testdata
	updatedItem1 := dto.ItemInput{
		Name:         Item1.Name,
		Price:        Item1.Price,
		Contributors: []uuid.UUID{User1.ID, User2.ID},
	}

	updatedItem2 := dto.ItemInput{
		Name:         Item2.Name,
		Price:        Item2.Price,
		Contributors: []uuid.UUID{User1.ID},
	}

	updatedBill := dto.BillUpdate{
		UpdatedAt: responseData.Data.UpdatedAt.Truncate(time.Second),
		Name:      "Updated Bill",
		Date:      Bill1.Date,
		Items:     []dto.ItemInput{updatedItem1, updatedItem2},
	}

	route := "/api/bill/"
	tests := []struct {
		description        string // description of the testcase case
		parameter          string
		inputData          dto.BillUpdate
		cookie             *http.Cookie // cookie of the testcase
		expectedCode       int          // expected HTTP status code
		expectedMessage    string       // expected message in response body
		expectReturn       bool         // expected return value
		expectReturnedData entity.Bill  // expected return
	}{
		{
			description:     "Test successful bill update",
			parameter:       Bill1.ID.String(),
			inputData:       updatedBill,
			cookie:          &http.Cookie{Name: sessionCookie, Value: CookieUser1.ID.String()},
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgBillUpdate,
			expectReturn:    true,
			expectReturnedData: entity.Bill{
				Base: entity.Base{
					ID: Bill1.ID,
				},
				Name:    Bill1.Name,
				OwnerID: User1.ID,
				Items:   []entity.Item{Item1, Item2},
			},
		},
	}

	for _, testcase := range tests {
		responseData, httpResponse, err := performBillRequest(http.MethodPut, route+testcase.parameter, testcase.inputData, testcase.cookie)

		// Assertion
		assert.NoError(t, err)
		assert.Equal(t, testcase.expectedCode, httpResponse.StatusCode)
		assert.Equal(t, testcase.expectedMessage, responseData.Message)
		assert.Equal(t, testcase.expectReturnedData.ID, responseData.Data.ID) // ID should not be changed
		assert.Equal(t, updatedBill.Name, responseData.Data.Name)
		assert.Equal(t, len(updatedBill.Items), len(responseData.Data.Items))
		for i, item := range responseData.Data.Items {
			// TODO: comment in to test if updated bill items do not change their IDs
			//assert.Equal(t, testcase.expectReturnedData.Items[i].ID, item.ID) // ID should not be changed
			assert.Equal(t, updatedBill.Items[i].Name, item.Name)
			assert.Equal(t, updatedBill.Items[i].Price, item.Price)
			assert.Equal(t, len(updatedBill.Items[i].Contributors), len(item.Contributors))
		}

	}
}
