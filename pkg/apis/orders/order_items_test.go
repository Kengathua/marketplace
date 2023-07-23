package orders

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/matawis/matawis/pkg/common"
	"github.com/matawis/matawis/pkg/models"
	"github.com/matawis/matawis/tests"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func Test_CustomerOrderItem(t *testing.T) {
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

	err := tx.Create(&customer).Error
	assert.Nil(err)

	customerCart := models.CustomerCart{
		CustomerID: *customer.ID,
		CartName:   "Cart One",
		CartCode:   "C-001",
	}

	err = tx.Create(&customerCart).Error
	assert.Nil(err)

	business_partner := models.BusinessPartner{Name: "Naivas Supermarket", BP_Code: "001"}
	err = tx.Create(&business_partner).Error
	assert.Nil(err)

	division := models.Division{DivisionName: "Appliances", DivisionCode: "D-001"}
	err = tx.Create(&division).Error
	assert.Nil(err)

	superCategory := models.SuperCategory{
		DivisionID: *division.ID, SuperCategoryName: "Large Appliances", SuperCategoryCode: "SC-001"}
	err = tx.Create(&superCategory).Error
	assert.Nil(err)

	category := models.Category{
		SuperCategoryID: *superCategory.ID, CategoryName: "Refrigerator", CategoryCode: "C-001"}
	err = tx.Create(&category).Error
	assert.Nil(err)

	itemType := models.ItemType{
		CategoryID: *category.ID,
		TypeName:   "Fridge",
		TypeCode:   "001",
	}
	err = tx.Create(&itemType).Error
	assert.Nil(err)

	brand := models.Brand{
		BrandName: "Samsung",
		BrandCode: "001",
	}
	err = tx.Create(&brand).Error
	assert.Nil(err)

	model := models.Model{
		BrandID:     *brand.ID,
		ItemTypeID:  *itemType.ID,
		ModelNumber: "S 10",
		ModelCode:   "001",
	}
	err = tx.Create(&model).Error
	assert.Nil(err)

	item := models.Item{
		ModelID: *model.ID, ItemName: "Samsung S-001 Fridge", ItemCode: "I-001",
	}
	err = tx.Create(&item).Error
	assert.Nil(err)

	catalogItem := models.CatalogItem{
		ItemID:         *item.ID,
		MarkedPrice:    decimal.NewFromFloat(13000),
		DiscountAmount: decimal.NewFromFloat(1000),
		SellingPrice:   decimal.NewFromFloat(10000),
		ThresholdPrice: decimal.NewFromFloat(9000),
	}
	err = tx.Create(&catalogItem).Error
	assert.Nil(err)

	customerCartItem := models.CustomerCartItem{
		CustomerCartID: *customerCart.ID,
		CatalogItemID:  *catalogItem.ID,
		UnitPrice:      "13000",
		Quantity:       "1",
		TotalPrice:     "13000",
	}

	err = tx.Create(&customerCartItem).Error
	assert.Nil(err)

	customerOrder := models.CustomerOrder{
		CustomerCartID: *customerCart.ID,
		OrderName:      "Order One",
		OrderCode:      "001",
	}
	err = tx.Create(&customerOrder).Error
	assert.Nil(err)

	customerOrderItem := models.CustomerOrderItem{
		CustomerOrderID:    *customerOrder.ID,
		CustomerCartItemID: *customerCartItem.ID,
		UnitPrice:          "13000",
		Quantity:           "1",
		TotalPrice:         "13000",
	}
	err = tx.Create(&customerOrderItem).Error
	assert.Nil(err)
	tx.Commit()

	testOrderItemID := *customerOrderItem.ID

	cartItemTests := []struct {
		description  string                 // description of the test case
		route        string                 // route path to test
		expectedCode int                    // expected HTTP status code
		httpMethod   string                 // http method to be tested
		payload      map[string]interface{} // payload expected for the test case
	}{
		{
			description:  "get HTTP status 200",
			route:        "/api/v1/orders/customer_order_items",
			expectedCode: http.StatusOK, // 200
			httpMethod:   "GET",
		},
		{
			description:  "get HTTP status 200",
			route:        fmt.Sprintf("/api/v1/orders/customer_order_items/%s", testOrderItemID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "GET",
		},
		{
			description:  "post HTTP status 201",
			route:        "/api/v1/orders/customer_order_items",
			expectedCode: http.StatusCreated, // 201
			httpMethod:   "POST",
			payload: map[string]interface{}{
				"customer_order_id": *customerOrder.ID,
				"customer_cart_id":  *customerCartItem.ID,
				"unit_price":       "12000",
				"quantity":         "1",
				"total_price":      "12000",
			},
		},
		{
			description:  "put HTTP status 201",
			route:        fmt.Sprintf("/api/v1/orders/customer_order_items/%s", testOrderItemID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "PUT",
			payload: map[string]interface{}{
				"customer_order_id": *customerOrder.ID,
				"customer_cart_id":  *customerCartItem.ID,
				"unit_price":       "12500",
				"quantity":         "1",
				"total_price":      "12500",
			},
		},
		{
			description:  "delete HTTP status 200",
			route:        fmt.Sprintf("/api/v1/orders/customer_order_items/%s", testOrderItemID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "DELETE",
		},
	}

	APITests(t, cartItemTests)
}
