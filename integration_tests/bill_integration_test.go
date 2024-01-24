package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"split-the-bill-server/presentation/dto"
	"testing"
)

type BillResponseDTO struct {
	Message string   `json:"message"`
	Data    dto.Bill `json:"data"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Start of GetBillByID test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestGetBillByID_Success(t *testing.T) {

	// Get Authentication Token
	cookieToken, setupErr := login(User1.Email, Password)
	if setupErr != nil {
		t.Fatalf("Error in test setup while logging in: %s", setupErr.Error())
	}

	route := "/api/bill/"
	cookie := &http.Cookie{Name: "session_cookie", Value: cookieToken}

	// Create http request
	req := httptest.NewRequest(http.MethodGet, route+Bill1.ID.String(), nil)
	// add cookie to request
	req.AddCookie(cookie)

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
	// Parse response body to GeneralResponseDTO
	var response BillResponseDTO
	if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
		t.Fatalf("Error while parsing response body: %s", err.Error())
	}
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEqual(t, uuid.Nil, response.Data.ID)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// End of GetBillByID test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Start of UpdateBill test cases
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func performUpdateBillRequest(testInputBill dto.Bill, username string) (BillResponseDTO, *http.Response, error) {
	route := "/api/bill/"

	// Get Authentication Token
	cookieToken, setupErr := login(User1.Email, Password)
	if setupErr != nil {
		return BillResponseDTO{}, nil, setupErr
	}

	inputJSON, _ := json.Marshal(testInputBill)
	cookie := &http.Cookie{Name: "session_cookie", Value: cookieToken}

	// Create http request
	req := httptest.NewRequest("PUT", route+testInputBill.ID.String(), bytes.NewReader(inputJSON))
	req.Header.Set("Content-Type", "application/json")
	// add cookie to request
	req.AddCookie(cookie)

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
		return BillResponseDTO{}, resp, err
	}
	// Parse response body to GeneralResponseDTO
	var response BillResponseDTO
	if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
		return BillResponseDTO{}, resp, err
	}
	return response, resp, nil
}

// TestUpdateBill_Core_Field_Success tests the update of bill fields.
func TestUpdateBill_Core_Field_Success(t *testing.T) {

	// create test data, update single core bill field
	testInputItem1 := dto.Item{
		ID: Item1.ID,
		BaseItem: dto.BaseItem{
			Name:           Item1.Name,
			Price:          Item1.Price,
			ContributorIDs: []uuid.UUID{User1.ID, User2.ID},
		}}

	testInputItem2 := dto.Item{
		ID: Item2.ID,
		BaseItem: dto.BaseItem{
			Name:           Item2.Name,
			Price:          Item2.Price,
			ContributorIDs: []uuid.UUID{User1.ID},
		},
	}

	testInputBill := dto.Bill{
		ID: Bill1.ID,
		BaseBill: dto.BaseBill{
			Name:    "Updated Bill",
			OwnerID: User1.ID,
			GroupID: Bill1.GroupID,
			Items:   []dto.Item{testInputItem1, testInputItem2}}}

	responseData, httpResponse, err := performUpdateBillRequest(testInputBill, User1.Email)
	if err != nil {
		t.Fatalf("Error during setup while performing request: %s", err.Error())
	}

	// get the stored bill from the storage for comparison
	storedBill, err := getStoredBill(Bill1.ID)

	assert.Nil(t, err)
	assert.Equal(t, 200, httpResponse.StatusCode)
	// validate response
	assert.Equal(t, testInputBill.Name, responseData.Data.Name)
	assert.Equal(t, testInputBill.ID, responseData.Data.ID) // ID should not be changed
	assert.Equal(t, testInputBill.OwnerID, responseData.Data.OwnerID)
	assert.Equal(t, testInputBill.GroupID, responseData.Data.GroupID)
	assert.Equal(t, len(testInputBill.Items), len(responseData.Data.Items))
	// validate updated bill in storage
	assert.Equal(t, testInputBill.Name, storedBill.Name)
	assert.Equal(t, testInputBill.ID, storedBill.ID) // ID should not be changed
	assert.Equal(t, len(testInputBill.Items), len(storedBill.Items))
	for i, item := range storedBill.Items {
		assert.Equal(t, testInputBill.Items[i].ID, item.ID) // ID should not be changed
		assert.Equal(t, testInputBill.Items[i].Name, item.Name)
		assert.Equal(t, testInputBill.Items[i].Price, item.Price)
		assert.Equal(t, len(testInputBill.Items[i].ContributorIDs), len(item.Contributors))
	}
}

// TestUpdateBill_Item_Field_Success tests the update of the items of a bill.
func TestUpdateBill_Item_Field_Success(t *testing.T) {

	// create test data, update item field
	testInputItem1 := dto.Item{
		ID: Item1.ID,
		BaseItem: dto.BaseItem{
			Name:           "Updated Item",
			Price:          Item1.Price,
			ContributorIDs: []uuid.UUID{User1.ID, User2.ID},
		},
	}

	testInputItem2 := dto.Item{
		ID: Item2.ID,
		BaseItem: dto.BaseItem{
			Name:           Item2.Name,
			Price:          Item2.Price,
			ContributorIDs: []uuid.UUID{User1.ID},
		},
	}

	testInputBill := dto.Bill{
		ID: Bill1.ID,
		BaseBill: dto.BaseBill{
			Name:    "Updated Bill",
			OwnerID: User1.ID,
			GroupID: Bill1.GroupID,
			Items:   []dto.Item{testInputItem1, testInputItem2},
		},
	}

	responseData, httpResponse, err := performUpdateBillRequest(testInputBill, User1.Email)
	if err != nil {
		t.Fatalf("Error during setup while performing request: %s", err.Error())
	}
	log.Println(responseData.Data.Items)

	// get the stored bill from the storage for comparison
	storedBill, err := getStoredBill(testInputBill.ID)

	log.Println(storedBill.Items)

	assert.Nil(t, err)
	assert.Equal(t, 200, httpResponse.StatusCode)
	// validate response
	assert.Equal(t, testInputBill.Name, responseData.Data.Name)
	assert.Equal(t, testInputBill.ID, responseData.Data.ID) // ID should not be changed
	assert.Equal(t, testInputBill.OwnerID, responseData.Data.OwnerID)
	assert.Equal(t, testInputBill.GroupID, responseData.Data.GroupID)
	assert.Equal(t, len(testInputBill.Items), len(responseData.Data.Items))
	// validate updated bill in storage
	assert.Equal(t, testInputBill.Name, storedBill.Name)
	assert.Equal(t, testInputBill.ID, storedBill.ID) // ID should not be changed
	assert.Equal(t, len(testInputBill.Items), len(storedBill.Items))
	for i, item := range storedBill.Items {
		assert.Equal(t, testInputBill.Items[i].ID, item.ID) // ID should not be changed
		assert.Equal(t, testInputBill.Items[i].Name, item.Name)
		assert.Equal(t, testInputBill.Items[i].Price, item.Price)
		assert.Equal(t, len(testInputBill.Items[i].ContributorIDs), len(item.Contributors))
	}

}

// TODO: decide on the test setup
/*func TestUpdateBill(t *testing.T) {

	// Get Authentication Token
	cookieToken, setupErr := login("felix@gmail.com", "test")
	if setupErr != nil {
		t.Fatalf("Error in test setup while logging in: %s", setupErr.Error())
	}

	route := "/api/bill/"
	tests := []struct {
		description        string // description of the testcase case
		parameter          string
		inputBill          dto.Bill
		inputJSON          []byte
		cookie             *http.Cookie // cookie of the testcase
		expectedCode       int          // expected HTTP status code
		expectedMessage    string       // expected message in response body
		expectReturn       bool         // expected return value
		expectReturnedData entity.Bill  // expected return
	}{
		{
			description: "Test successful: bill update",
			parameter:   Bill1.ID.String(),
			inputBill: dto.Bill{
				ID: Bill1.ID,
				BaseBill: dto.BaseBill{
					Name:    "Updated Bill",
					OwnerID: User1.ID,
					Items: []dto.Item{
						{
							ID: Item1.ID,
							BaseItem: dto.BaseItem{
								Name:           Item1.Name,
								Price:          Item1.Price,
								ContributorIDs: []uuid.UUID{User1.ID, User2.ID},
							},
						},
						{
							ID: Item2.ID,
							BaseItem: dto.BaseItem{
								Name:           Item2.Name,
								Price:          Item2.Price,
								ContributorIDs: []uuid.UUID{User1.ID},
							},
						},
					},
				},
			},
			cookie:          &http.Cookie{Name: "session_cookie", Value: cookieToken},
			expectedCode:    200,
			expectedMessage: handler.SuccessMsgBillUpdate,
			expectReturn:    true,
			expectReturnedData: entity.Bill{ // TODO: rework expectReturnedData
				Base: entity.Base{
					ID: Bill1.ID,
				},
				Name:    Bill1.Name,
				OwnerID: User1.ID,
				Items:   []entity.Item{Item1, Item2},
			},
		},
		{
			description: "Test unsuccessful: item id missing",
			parameter:   Bill1.ID.String(),
			inputBill: dto.Bill{
				ID: Bill1.ID,
				BaseBill: dto.BaseBill{
					Name:    "Updated Bill",
					OwnerID: User1.ID,
					Items: []dto.Item{
						{
							BaseItem: dto.BaseItem{
								Name:           Item1.Name,
								Price:          Item1.Price,
								ContributorIDs: []uuid.UUID{User1.ID, User2.ID},
							},
						},
						{
							BaseItem: dto.BaseItem{
								Name:           Item2.Name,
								Price:          Item2.Price,
								ContributorIDs: []uuid.UUID{User1.ID},
							},
						},
					},
				},
			},
			cookie:             &http.Cookie{Name: "session_cookie", Value: cookieToken},
			expectedCode:       500,
			expectedMessage:    fmt.Sprintf(handler.ErrMsgBillUpdate, storage.NoSuchItemError),
			expectReturn:       false,
			expectReturnedData: entity.Bill{},
		},
	}

	for _, testcase := range tests {
		inputJson, _ := json.Marshal(testcase.inputBill)
		// Create http request
		req := httptest.NewRequest("PUT", route+testcase.parameter, bytes.NewReader(inputJson))
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
		// Parse response body to GeneralResponseDTO
		var response BillResponseDTO
		if err = json.Unmarshal(body, &response); err != nil { // Parse []byte to go struct pointer
			t.Fatalf("Error while parsing response body: %s", err.Error())
		}

		// Assertion
		assert.NoError(t, err)
		assert.Equal(t, testcase.expectedCode, resp.StatusCode)
		assert.Equal(t, testcase.expectedMessage, response.Message)
		if testcase.expectReturn {
			assert.Equal(t, testcase.expectReturnedData.ID, response.Data.ID) // ID should not be changed
			assert.Equal(t, "Updated Bill", response.Data.Name)
			assert.Equal(t, len(Bill1.Items), len(response.Data.Items))
			for i, item := range response.Data.Items {
				assert.Equal(t, Bill1.Items[i].ID, item.ID) // ID should not be changed
				assert.Equal(t, Bill1.Items[i].Name, item.Name)
				assert.Equal(t, Bill1.Items[i].Price, item.Price)
				assert.Equal(t, len(Bill1.Items[i].Contributors), len(item.ContributorIDs))
			}
		}

	}
}*/
