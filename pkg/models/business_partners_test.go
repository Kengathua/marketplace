package models

import (
	"testing"

	"github.com/Kengathua/marketplace/tests"
	"github.com/stretchr/testify/assert"
)

func TestBusinessPartner(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	business_partner := BusinessPartner{
		Name:           "Quick Shop",
		BP_Code:        "001",
		MainBranchCode: "",
		Description:    "Quick Shop",
	}

	t.Run("test create business partner", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&business_partner).Error
		assert.Nil(err)

		count := int64(0)
		result := &BusinessPartner{}
		db.First(result).Count(&count)
		assert.Equal(*business_partner.ID, *result.ID)
	})
}
