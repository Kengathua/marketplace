package models

import (
	"fmt"
	"testing"

	"github.com/Kengathua/marketplace/pkg/common"
	"github.com/Kengathua/marketplace/tests"
	"github.com/stretchr/testify/assert"
)

func TestCustomerCart(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	customer := Customer{
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

	customerCart := CustomerCart{
		CartName: "Cart One",
		CartCode: "C-001",
	}

	t.Run("test create customer cart", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&customer).Error
		assert.Nil(err)

		fmt.Println("The customer id is ", *customer.ID)
		customerCart.CustomerID = *customer.ID
		err = db.Create(&customerCart).Error
		assert.Nil(err)

		count := int64(0)
		result := &CustomerCart{}
		db.First(result).Count(&count)
		assert.Equal(*customerCart.ID, *result.ID)
	})
}
