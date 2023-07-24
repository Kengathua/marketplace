package retail

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/Kengathua/marketplace/pkg/models"
	"github.com/Kengathua/marketplace/tests"
	"github.com/stretchr/testify/assert"
)

func Test_Brand_item_types(t *testing.T) {
	db := tests.GetTestDB()
	assert := assert.New(t)
	tx := db.Begin()
	defer tests.TruncateDatabase(db)

	business_partner := models.BusinessPartner{Name: "Naivas Supermarket", BP_Code: "001"}
	err := tx.Create(&business_partner).Error
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

	brandItemType := models.BrandItemType{
		BrandID:    *brand.ID,
		ItemTypeID: *itemType.ID,
	}
	err = tx.Create(&brandItemType).Error
	assert.Nil(err)

	tx.Commit()

	brandItemTypeID := *brandItemType.ID

	brandItemTypeTests := []struct {
		description  string                 // description of the test case
		route        string                 // route path to test
		expectedCode int                    // expected HTTP status code
		httpMethod   string                 // http method to be tested
		payload      map[string]interface{} // payload expected for the test case
	}{
		{
			description:  "get HTTP status 200",
			route:        "/api/v1/retail/brand_item_types",
			expectedCode: http.StatusOK, // 200
			httpMethod:   "GET",
		},
		{
			description:  "get HTTP status 200",
			route:        fmt.Sprintf("/api/v1/retail/brand_item_types/%s", brandItemTypeID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "GET",
		},
		{
			description:  "post HTTP status 201",
			route:        "/api/v1/retail/brand_item_types",
			expectedCode: http.StatusCreated, // 201
			httpMethod:   "POST",
			payload: map[string]interface{}{
				"brand_id":     *brand.ID,
				"item_type_id": *itemType.ID,
			},
		},
		{
			description:  "put HTTP status 201",
			route:        fmt.Sprintf("/api/v1/retail/brand_item_types/%s", brandItemTypeID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "PUT",
			payload: map[string]interface{}{
				"brand_id":     *brand.ID,
				"item_type_id": *itemType.ID,
			},
		},
		{
			description:  "delete HTTP status 200",
			route:        fmt.Sprintf("/api/v1/retail/brand_item_types/%s", brandItemTypeID),
			expectedCode: http.StatusOK, // 200
			httpMethod:   "DELETE",
		},
	}

	APITests(t, brandItemTypeTests)
}
