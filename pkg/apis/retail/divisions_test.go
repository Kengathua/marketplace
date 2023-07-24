package retail

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kengathua/marketplace/pkg/models"
	"github.com/Kengathua/marketplace/tests"
	"github.com/stretchr/testify/assert"
)

func Test_Division(t *testing.T) {
	db := tests.GetTestDB()
	assert := assert.New(t)
	tx := db.Begin()
	defer tests.TruncateDatabase(db)

	division := models.Division{DivisionName: "Appliances", DivisionCode: "D-001"}
	err := tx.Create(&division).Error
	assert.Nil(err)

	tx.Commit()

	testdivisionID := *division.ID

	divisionTests := []struct {
		description  string                 // description of the test case
		route        string                 // route path to test
		expectedCode int                    // expected HTTP status code
		httpMethod   string                 // http method to be tested
		payload      map[string]interface{} // payload expected for the test case
	}{
		{
			description:  "get HTTP status 200",
			route:        "/api/v1/retail/divisions",
			expectedCode: http.StatusOK, // 200
			httpMethod:   "GET",
		},
		{
			description:  "get HTTP status 200",
			route:        fmt.Sprintf("/api/v1/retail/divisions/%s", testdivisionID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "GET",
		},
		{
			description:  "post HTTP status 201",
			route:        "/api/v1/retail/divisions",
			expectedCode: http.StatusCreated, // 201
			httpMethod:   "POST",
			payload: map[string]interface{}{
				"division_name": "Electronics & Appliances",
				"division_code": "Cart One",
				"description":   "001",
			},
		},
		{
			description:  "put HTTP status 201",
			route:        fmt.Sprintf("/api/v1/retail/divisions/%s", testdivisionID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "PUT",
			payload: map[string]interface{}{
				"division_name": "Electricals & Appliances",
				"division_code": "Cart One Update",
				"description":   "0001",
			},
		},
		{
			description:  "delete HTTP status 200",
			route:        fmt.Sprintf("/api/v1/retail/divisions/%s", testdivisionID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "DELETE",
		},
	}

	APITests(t, divisionTests)
}
