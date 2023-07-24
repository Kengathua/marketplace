package models

import (
	"testing"

	"github.com/Kengathua/marketplace/pkg/common"
	"github.com/Kengathua/marketplace/tests"
	"github.com/stretchr/testify/assert"
)

func TestCustomer(t *testing.T) {
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

	t.Run("test create customer", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&customer).Error
		assert.Nil(err)

		count := int64(0)
		result := &Customer{}
		db.First(result).Count(&count)
		assert.Equal(*customer.ID, *result.ID)
	})
}
