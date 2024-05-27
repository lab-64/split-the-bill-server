package integration_tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"split-the-bill-server/presentation/dto"
	"split-the-bill-server/presentation/handler"
	"split-the-bill-server/storage/database/entity"
	"testing"
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

	// Testdata
	updatedItem1 := dto.ItemInput{
		Name:         &Item1.Name,
		Price:        &Item1.Price,
		Contributors: &dto.Changes[uuid.UUID]{Add: &[]uuid.UUID{User1.ID, User2.ID}},
	}

	updatedItem2 := dto.ItemInput{
		Name:         &Item2.Name,
		Price:        &Item2.Price,
		Contributors: &dto.Changes[uuid.UUID]{Add: &[]uuid.UUID{User1.ID}},
	}

	updatedName := "Updated Bill"
	updatedBill := dto.BillUpdate{
		Name:  &updatedName,
		Date:  &Bill1.Date,
		Items: &dto.Changes[dto.ItemInput]{Update: &[]dto.ItemInput{updatedItem1, updatedItem2}},
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
				Name:    "Updated Bill",
				OwnerID: User1.ID,
				Items:   []entity.Item{Item1, Item2},
			},
		},
		{
			description: "Test update only bill name",
			parameter:   Bill1.ID.String(),
			inputData: dto.BillUpdate{
				Name: strToPtr("New Updated Bill"),
			},
			cookie:          &http.Cookie{Name: sessionCookie, Value: CookieUser1.ID.String()},
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgBillUpdate,
			expectReturn:    true,
			expectReturnedData: entity.Bill{
				Base: entity.Base{
					ID: Bill1.ID,
				},
				Name:    "New Updated Bill",
				OwnerID: Bill1.OwnerID,
				Items:   Bill1.Items,
			},
		},
		{
			description: "Test remove item from bill",
			parameter:   Bill2.ID.String(),
			inputData: dto.BillUpdate{
				Items: &dto.Changes[dto.ItemInput]{Remove: &[]dto.ItemInput{{ID: Item3.ID}}},
			},
			cookie:          &http.Cookie{Name: sessionCookie, Value: CookieUser1.ID.String()},
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgBillUpdate,
			expectReturn:    true,
			expectReturnedData: entity.Bill{
				Base: entity.Base{
					ID: Bill2.ID,
				},
				Name:    Bill2.Name,
				OwnerID: Bill2.OwnerID,
				Items:   []entity.Item{},
			},
		},
		{
			description: "Test add item to bill",
			parameter:   Bill3.ID.String(),
			inputData: dto.BillUpdate{
				Items: &dto.Changes[dto.ItemInput]{Add: &[]dto.ItemInput{{Name: &Item3.Name, Price: &Item3.Price, Contributors: &dto.Changes[uuid.UUID]{Add: &[]uuid.UUID{User1.ID}}}}},
			},
			cookie:          &http.Cookie{Name: sessionCookie, Value: CookieUser1.ID.String()},
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgBillUpdate,
			expectReturn:    true,
			expectReturnedData: entity.Bill{
				Base: entity.Base{
					ID: Bill3.ID,
				},
				Name:    Bill3.Name,
				OwnerID: Bill3.OwnerID,
				Items: append(Bill3.Items, entity.Item{
					Base:         entity.Base{ID: Item3.ID},
					Name:         Item3.Name,
					Price:        Item3.Price,
					Contributors: []*entity.User{&User1},
				}),
			},
		},
	}

	for _, testcase := range tests {
		responseData, httpResponse, err := performBillRequest(http.MethodPut, route+testcase.parameter, testcase.inputData, testcase.cookie)

		// Assertion
		assert.NoErrorf(t, err, testcase.description)
		assert.Equalf(t, testcase.expectedCode, httpResponse.StatusCode, testcase.description)
		assert.Equalf(t, testcase.expectedMessage, responseData.Message, testcase.description)
		assert.Equalf(t, testcase.expectReturnedData.ID, responseData.Data.ID, testcase.description) // ID should not be changed
		assert.Equalf(t, testcase.expectReturnedData.Name, responseData.Data.Name, testcase.description)
		log.Println(responseData.Data.Items)
		assert.Equalf(t, len(testcase.expectReturnedData.Items), len(responseData.Data.Items), testcase.description)
		for i, responseItem := range responseData.Data.Items {
			assert.Equalf(t, testcase.expectReturnedData.Items[i].Name, responseItem.Name, testcase.description)
			assert.Equalf(t, testcase.expectReturnedData.Items[i].Price, responseItem.Price, testcase.description)
			assert.Equalf(t, len(testcase.expectReturnedData.Items[i].Contributors), len(responseItem.Contributors), testcase.description)
			// if a new item is added, we cannot compare the ID
			if testcase.description != "Test add item to bill" {
				assert.Equalf(t, testcase.expectReturnedData.Items[i].ID, responseItem.ID, testcase.description) // itemID should not be changed
			}
		}
	}
}

func strToPtr(s string) *string {
	return &s
}
