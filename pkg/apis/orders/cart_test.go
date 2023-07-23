package orders

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/matawis/matawis/pkg/common"
	"github.com/matawis/matawis/pkg/models"
	"github.com/matawis/matawis/tests"
	"github.com/stretchr/testify/assert"
)

func Test_CustomerCart(t *testing.T) {
	db := tests.GetTestDB()
	assert := assert.New(t)
	tx := db.Begin()
	defer tests.TruncateDatabase(db)

	customer := models.Customer{
		BioData: common.BioData{
			Title:       "Mr",
			FirstName:   "Test",
			LastName:    "User",
			OtherNames:  "Development",
			Email:       "testuser@email.com",
			PhoneNumber: "254712345678",
			Gender:      "MALE",
		},
		CustomerNumber: "C-001",
	}

	customerCart := models.CustomerCart{
		CartName: "Cart One",
		CartCode: "C-001",
	}

	err := tx.Create(&customer).Error
	assert.Nil(err)

	customerCart.CustomerID = *customer.ID
	err = tx.Create(&customerCart).Error
	assert.Nil(err)
	tx.Commit()

	testCartID := *customerCart.ID

	cartTests := []struct {
		description  string                 // description of the test case
		route        string                 // route path to test
		expectedCode int                    // expected HTTP status code
		httpMethod   string                 // http method to be tested
		payload      map[string]interface{} // payload expected for the test case
	}{
		{
			description:  "get HTTP status 200",
			route:        "/api/v1/orders/customer_carts",
			expectedCode: http.StatusOK, // 200
			httpMethod:   "GET",
		},
		{
			description:  "get HTTP status 200",
			route:        fmt.Sprintf("/api/v1/orders/customer_carts/%s", testCartID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "GET",
		},
		{
			description:  "post HTTP status 201",
			route:        "/api/v1/orders/customer_carts",
			expectedCode: http.StatusCreated, // 201
			httpMethod:   "POST",
			payload: map[string]interface{}{
				"customer_id": *customer.ID,
				"cart_name":   "Cart One",
				"cart_code":   "001",
			},
		},
		{
			description:  "put HTTP status 201",
			route:        fmt.Sprintf("/api/v1/orders/customer_carts/%s", testCartID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "PUT",
			payload: map[string]interface{}{
				"customer_id": *customer.ID,
				"cart_name":   "Cart One Update",
				"cart_code":   "0001",
			},
		},
		{
			description:  "delete HTTP status 200",
			route:        fmt.Sprintf("/api/v1/orders/customer_carts/%s", testCartID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "DELETE",
		},
	}

	APITests(t, cartTests)
}
