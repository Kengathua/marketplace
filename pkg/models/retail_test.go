package models

import (
	"testing"

	"github.com/Kengathua/marketplace/tests"
	"github.com/stretchr/testify/assert"
)

func TestDivision(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	division := Division{
		DivisionName: "Acessories & Electronics",
		DivisionCode: "D-001",
		Description:  "Acessories & Electronics",
	}

	t.Run("test create division", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&division).Error
		assert.Nil(err)

		count := int64(0)
		result := &Division{}
		db.First(result).Count(&count)
		assert.Equal(*division.ID, *result.ID)
	})
}

func TestSuperCategory(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	division := Division{
		DivisionName: "Phones & Tablets",
		DivisionCode: "D-001",
		Description:  "Phones & Tablets",
	}

	supCategory := SuperCategory{
		SuperCategoryName: "Mobile Phones",
		SuperCategoryCode: "SC-001",
		Description:       "Mobile Phones",
	}

	t.Run("test create super category", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&division).Error
		assert.Nil(err)

		supCategory.DivisionID = *division.ID
		err = db.Create(&supCategory).Error
		assert.Nil(err)

		count := int64(0)
		result := &SuperCategory{}
		db.First(result).Count(&count)
		assert.Equal(*supCategory.ID, *result.ID)
	})
}

func TestCategory(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	division := Division{
		DivisionName: "Phones & Tablets",
		DivisionCode: "D-001",
		Description:  "Phones & Tablets",
	}

	supCategory := SuperCategory{
		SuperCategoryName: "Mobile Phones",
		SuperCategoryCode: "SC-001",
		Description:       "Mobile Phones",
	}

	category := Category{
		CategoryName: "Smart Phones",
		CategoryCode: "C-001",
		Description:  "Smart Phones",
	}

	t.Run("test create category", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&division).Error
		assert.Nil(err)

		supCategory.DivisionID = *division.ID
		err = db.Create(&supCategory).Error
		assert.Nil(err)

		category.SuperCategoryID = *supCategory.ID
		err = db.Create(&category).Error
		assert.Nil(err)

		count := int64(0)
		result := &Category{}
		db.First(result).Count(&count)
		assert.Equal(*category.ID, *result.ID)
	})
}

func TestItemType(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	division := Division{
		DivisionName: "Phones & Tablets",
		DivisionCode: "D-001",
		Description:  "Phones & Tablets",
	}

	supCategory := SuperCategory{
		SuperCategoryName: "Mobile Phones",
		SuperCategoryCode: "SC-001",
		Description:       "Mobile Phones",
	}

	category := Category{
		CategoryName: "Smart Phones",
		CategoryCode: "C-001",
		Description:  "Smart Phones",
	}

	itemType := ItemType{
		TypeName: "Phone",
		TypeCode: "IT-001",
	}

	t.Run("test create item type", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&division).Error
		assert.Nil(err)

		supCategory.DivisionID = *division.ID
		err = db.Create(&supCategory).Error
		assert.Nil(err)

		category.SuperCategoryID = *supCategory.ID
		err = db.Create(&category).Error
		assert.Nil(err)

		itemType.CategoryID = *category.ID
		err = db.Create(&itemType).Error
		assert.Nil(err)

		count := int64(0)
		result := &ItemType{}
		db.First(result).Count(&count)
		assert.Equal(*itemType.ID, *result.ID)
	})
}

func TestBrand(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	division := Division{
		DivisionName: "Phones & Tablets",
		DivisionCode: "D-001",
		Description:  "Phones & Tablets",
	}

	supCategory := SuperCategory{
		SuperCategoryName: "Mobile Phones",
		SuperCategoryCode: "SC-001",
		Description:       "Mobile Phones",
	}

	category := Category{
		CategoryName: "Smart Phones",
		CategoryCode: "C-001",
		Description:  "Smart Phones",
	}

	itemType := ItemType{
		TypeName: "Phone",
		TypeCode: "IT-001",
	}

	brand := Brand{
		BrandName: "Samsung",
		BrandCode: "B-001",
	}

	t.Run("test create brand", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&division).Error
		assert.Nil(err)

		supCategory.DivisionID = *division.ID
		err = db.Create(&supCategory).Error
		assert.Nil(err)

		category.SuperCategoryID = *supCategory.ID
		err = db.Create(&category).Error
		assert.Nil(err)

		itemType.CategoryID = *category.ID
		err = db.Create(&itemType).Error
		assert.Nil(err)

		err = db.Create(&brand).Error
		assert.Nil(err)

		count := int64(0)
		result := &Brand{}
		db.First(result).Count(&count)
		assert.Equal(*brand.ID, *result.ID)
	})
}

func TestBrandItemType(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	division := Division{
		DivisionName: "Phones & Tablets",
		DivisionCode: "D-001",
		Description:  "Phones & Tablets",
	}

	supCategory := SuperCategory{
		SuperCategoryName: "Mobile Phones",
		SuperCategoryCode: "SC-001",
		Description:       "Mobile Phones",
	}

	category := Category{
		CategoryName: "Smart Phones",
		CategoryCode: "C-001",
		Description:  "Smart Phones",
	}

	itemType := ItemType{
		TypeName: "Phone",
		TypeCode: "IT-001",
	}

	brand := Brand{
		BrandName: "Samsung",
		BrandCode: "B-001",
	}

	brandItemType := BrandItemType{}

	t.Run("test create brand item type", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&division).Error
		assert.Nil(err)

		supCategory.DivisionID = *division.ID
		err = db.Create(&supCategory).Error
		assert.Nil(err)

		category.SuperCategoryID = *supCategory.ID
		err = db.Create(&category).Error
		assert.Nil(err)

		itemType.CategoryID = *category.ID
		err = db.Create(&itemType).Error
		assert.Nil(err)

		err = db.Create(&brand).Error
		assert.Nil(err)

		brandItemType.BrandID = *brand.ID
		brandItemType.ItemTypeID = *itemType.ID

		err = db.Create(&brandItemType).Error
		assert.Nil(err)

		count := int64(0)
		result := &BrandItemType{}
		db.First(result).Count(&count)
		assert.Equal(*brandItemType.ID, *result.ID)
	})
}

func TestModel(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	division := Division{
		DivisionName: "Phones & Tablets",
		DivisionCode: "D-001",
		Description:  "Phones & Tablets",
	}

	supCategory := SuperCategory{
		SuperCategoryName: "Mobile Phones",
		SuperCategoryCode: "SC-001",
		Description:       "Mobile Phones",
	}

	category := Category{
		CategoryName: "Smart Phones",
		CategoryCode: "C-001",
		Description:  "Smart Phones",
	}

	itemType := ItemType{
		TypeName: "Phone",
		TypeCode: "IT-001",
	}

	brand := Brand{
		BrandName: "Samsung",
		BrandCode: "B-001",
	}

	brandItemType := BrandItemType{}

	model := Model{
		ModelNumber: "S 10",
		ModelCode:   "M-001",
	}

	t.Run("test create model", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&division).Error
		assert.Nil(err)

		supCategory.DivisionID = *division.ID
		err = db.Create(&supCategory).Error
		assert.Nil(err)

		category.SuperCategoryID = *supCategory.ID
		err = db.Create(&category).Error
		assert.Nil(err)

		itemType.CategoryID = *category.ID
		err = db.Create(&itemType).Error
		assert.Nil(err)

		err = db.Create(&brand).Error
		assert.Nil(err)

		brandItemType.BrandID = *brand.ID
		brandItemType.ItemTypeID = *itemType.ID
		err = db.Create(&brandItemType).Error
		assert.Nil(err)

		model.BrandID = *brand.ID
		model.ItemTypeID = *itemType.ID
		err = db.Create(&model).Error
		assert.Nil(err)

		count := int64(0)
		result := &Model{}
		db.First(result).Count(&count)
		assert.Equal(*model.ID, *result.ID)
	})
}

func TestItem(t *testing.T) {
	db := TestDB
	assert := assert.New(t)
	defer tests.TruncateDatabase(db)

	division := Division{
		DivisionName: "Phones & Tablets",
		DivisionCode: "D-001",
		Description:  "Phones & Tablets",
	}

	supCategory := SuperCategory{
		SuperCategoryName: "Mobile Phones",
		SuperCategoryCode: "SC-001",
		Description:       "Mobile Phones",
	}

	category := Category{
		CategoryName: "Smart Phones",
		CategoryCode: "C-001",
		Description:  "Smart Phones",
	}

	itemType := ItemType{
		TypeName: "Phone",
		TypeCode: "IT-001",
	}

	brand := Brand{
		BrandName: "Samsung",
		BrandCode: "B-001",
	}

	brandItemType := BrandItemType{}

	model := Model{
		ModelNumber: "S 10",
		ModelCode:   "M-001",
	}

	item := Item{
		ItemName: "Samsung S 10 Phone",
		ItemSize: "6.5 Inches",
		Barcode:  "76543234598",
		ItemCode: "I-001",
		MakeYear: "2022",
	}

	t.Run("test create item", func(t *testing.T) {
		defer tests.TruncateDatabase(db)

		err := db.Create(&division).Error
		assert.Nil(err)

		supCategory.DivisionID = *division.ID
		err = db.Create(&supCategory).Error
		assert.Nil(err)

		category.SuperCategoryID = *supCategory.ID
		err = db.Create(&category).Error
		assert.Nil(err)

		itemType.CategoryID = *category.ID
		err = db.Create(&itemType).Error
		assert.Nil(err)

		err = db.Create(&brand).Error
		assert.Nil(err)

		brandItemType.BrandID = *brand.ID
		brandItemType.ItemTypeID = *itemType.ID
		err = db.Create(&brandItemType).Error
		assert.Nil(err)

		model.BrandID = *brand.ID
		model.ItemTypeID = *itemType.ID
		err = db.Create(&model).Error
		assert.Nil(err)

		item.ModelID = *model.ID
		err = db.Create(&item).Error
		assert.Nil(err)

		count := int64(0)
		result := &Item{}
		db.First(result).Count(&count)
		assert.Equal(*item.ID, *result.ID)
	})
}
