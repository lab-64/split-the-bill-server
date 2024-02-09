package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
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

func TestUpdateBill(t *testing.T) {

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

	updatedBill := dto.BillInput{
		Name:    "Updated Bill",
		OwnerID: User1.ID,
		Items:   []dto.ItemInput{updatedItem1, updatedItem2},
	}

	inputJson, _ := json.Marshal(updatedBill)

	route := "/api/bill/"
	tests := []struct {
		description        string // description of the testcase case
		parameter          string
		inputJSON          []byte
		cookie             *http.Cookie // cookie of the testcase
		expectedCode       int          // expected HTTP status code
		expectedMessage    string       // expected message in response body
		expectReturn       bool         // expected return value
		expectReturnedData entity.Bill  // expected return
	}{
		{
			description:     "Test successful bill update",
			parameter:       Bill1.ID.String(),
			inputJSON:       inputJson,
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
		// Create http request
		req := httptest.NewRequest("PUT", route+testcase.parameter, bytes.NewReader(testcase.inputJSON))
		req.Header.Set("Content-Type", "application/json")
		// add cookie to request
		req.AddCookie(testcase.cookie)

		// Perform request
		resp, err := app.Test(req, -1)
		if resp != nil {
			defer resp.Body.Close()
		}
		if err != nil {
			t.Fatalf("Error while performing request: %s", err.Error())
		}

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Error while reading response body: %s", err.Error())
		}
		// Parse response body to GeneralResponse
		var response BillResponseDTO
		if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
			t.Fatalf("Error while parsing response body: %s", err.Error())
		}

		// Assertion
		assert.NoError(t, err)
		assert.Equal(t, testcase.expectedCode, resp.StatusCode)
		assert.Equal(t, testcase.expectedMessage, response.Message)
		assert.Equal(t, testcase.expectReturnedData.ID, response.Data.ID) // ID should not be changed
		assert.Equal(t, updatedBill.Name, response.Data.Name)
		assert.Equal(t, len(updatedBill.Items), len(response.Data.Items))
		for i, item := range response.Data.Items {
			// TODO: comment in to test if new implementation works
			//assert.Equal(t, testcase.expectReturnedData.Items[0].ID, item.ID) // ID should not be changed
			assert.Equal(t, updatedBill.Items[i].Name, item.Name)
			assert.Equal(t, updatedBill.Items[i].Price, item.Price)
			assert.Equal(t, len(updatedBill.Items[i].Contributors), len(item.Contributors))
		}

	}
}
