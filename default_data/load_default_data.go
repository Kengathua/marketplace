package main

import (
	"fmt"

	"github.com/shopspring/decimal"

	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/common"
	"github.com/matawis/matawis/pkg/config"
	"github.com/matawis/matawis/pkg/models"
)

func init() {
	config.ConnectToDb()
}

func main() {
	business_partner := models.BusinessPartner{Name: "Naivas Supermarket", BP_Code: "001"}
	config.DB.Find(&business_partner, "name = ?", "Naivas Supermarket")
	if business_partner.ID == uuid.Nil {
		config.DB.Create(&business_partner)
	} else {
		fmt.Printf("Business Partner %s already exists\n", business_partner.Name)
	}
	baseModel := common.BaseModel{BusinessPartner: business_partner.BP_Code}

	division1 := models.Division{DivisionName: "Appliances", DivisionCode: "D-001"}
	config.DB.Find(&division1, "division_name = ?", division1.DivisionName)
	if division1.ID == uuid.Nil {
		config.DB.Create(&division1)
	} else {
		fmt.Printf("Division %s already exists\n", division1.DivisionName)
	}

	division2 := models.Division{DivisionName: "Food & Beverages", DivisionCode: "D-002"}
	config.DB.Find(&division2, "division_name = ?", division2.DivisionName)
	if division2.ID == uuid.Nil {
		config.DB.Create(&division2)
	} else {
		fmt.Printf("Division %s already exists\n", division2.DivisionName)
	}

	superCategory1 := models.SuperCategory{
		DivisionID: division1.ID, SuperCategoryName: "Large Appliances", SuperCategoryCode: "SC-001"}
	config.DB.Find(&superCategory1, "super_category_name = ?", superCategory1.SuperCategoryName)
	if superCategory1.ID == uuid.Nil {
		config.DB.Create(&superCategory1)
	} else {
		fmt.Printf("Super category %s already exists\n", superCategory1.SuperCategoryName)
	}

	// https://www.betterhealth.vic.gov.au/health/healthyliving/fruit-and-vegetables
	superCategory2 := models.SuperCategory{
		DivisionID: division2.ID, SuperCategoryName: "Fruits & Vegetables", SuperCategoryCode: "SC-002"}
	config.DB.Find(&superCategory2, "super_category_name = ?", superCategory2.SuperCategoryName)
	if superCategory2.ID == uuid.Nil {
		config.DB.Create(&superCategory2)
	} else {
		fmt.Printf("Super category %s already exists\n", superCategory2.SuperCategoryName)
	}

	category1 := models.Category{
		SuperCategoryID: superCategory1.ID, CategoryName: "Refrigerator", CategoryCode: "C-001"}
	config.DB.Find(&category1, "category_name = ?", category1.CategoryName)
	if category1.ID == uuid.Nil {
		config.DB.Create(&category1)
	} else {
		fmt.Printf("Category %s already exists\n", category1.CategoryName)
	}

	category2 := models.Category{
		SuperCategoryID: superCategory2.ID, CategoryName: "Leafy Green", CategoryCode: "C-002"}
	config.DB.Find(&category2, "category_name = ?", category2.CategoryName)
	if category2.ID == uuid.Nil {
		config.DB.Create(&category2)
	} else {
		fmt.Printf("Category %s already exists\n", category2.CategoryName)
	}

	category3 := models.Category{
		SuperCategoryID: superCategory2.ID, CategoryName: "Cruciferous", CategoryCode: "C-003"}
	config.DB.Find(&category3, "category_name = ?", category3.CategoryName)
	if category3.ID == uuid.Nil {
		config.DB.Create(&category3)
	} else {
		fmt.Printf("Category %s already exists\n", category3.CategoryName)
	}

	category4 := models.Category{
		SuperCategoryID: superCategory2.ID, CategoryName: "Edible Stem", CategoryCode: "C-004"}
	config.DB.Find(&category4, "category_name = ?", category4.CategoryName)
	if category4.ID == uuid.Nil {
		config.DB.Create(&category4)
	} else {
		fmt.Printf("Category %s already exists\n", category4.CategoryName)
	}

	category5 := models.Category{
		SuperCategoryID: superCategory2.ID, CategoryName: "Root", CategoryCode: "C-005"}
	config.DB.Find(&category5, "category_name = ?", category5.CategoryName)
	if category5.ID == uuid.Nil {
		config.DB.Create(&category5)
	} else {
		fmt.Printf("Category %s already exists\n", category5.CategoryName)
	}

	category6 := models.Category{
		SuperCategoryID: superCategory2.ID, CategoryName: "Allium", CategoryCode: "C-006"}
	config.DB.Find(&category6, "category_name = ?", category6.CategoryName)
	if category6.ID == uuid.Nil {
		config.DB.Create(&category6)
	} else {
		fmt.Printf("Category %s already exists\n", category6.CategoryName)
	}

	category7 := models.Category{
		SuperCategoryID: superCategory2.ID, CategoryName: "Marrow", CategoryCode: "C-007"}
	config.DB.Find(&category7, "category_name = ?", category7.CategoryName)
	if category7.ID == uuid.Nil {
		config.DB.Create(&category7)
	} else {
		fmt.Printf("Category %s already exists\n", category7.CategoryName)
	}

	itemType := models.ItemType{
		CategoryID: category1.ID, TypeName: "Fridge", TypeCode: "D-001", BaseModel: baseModel}
	config.DB.Find(&itemType, "type_name = ?", itemType.TypeName)
	if itemType.ID == uuid.Nil {
		config.DB.Create(&itemType)
	} else {
		fmt.Printf("Item type %s already exists\n", itemType.TypeName)
	}

	brand := models.Brand{
		BrandName: "Samsung", BrandCode: "D-001", BaseModel: baseModel}
	config.DB.Find(&brand, "brand_name = ?", brand.BrandName)
	if brand.ID == uuid.Nil {
		config.DB.Create(&brand)
	} else {
		fmt.Printf("Brand %s already exists\n", brand.BrandName)
	}

	model := models.Model{
		BrandID: brand.ID, ModelNumber: "S-001", ModelCode: "M-001", BaseModel: baseModel}
	config.DB.Find(&model, "model_number = ?", model.ModelNumber)
	if model.ID == uuid.Nil {
		config.DB.Create(&model)
	} else {
		fmt.Printf("Model %s already exists\n", model.ModelNumber)
	}

	item := models.Item{
		ModelID: model.ID, ItemName: "Samsung S-001 Fridge", ItemCode: "I-001", BaseModel: baseModel}
	config.DB.Find(&item, "item_name = ?", item.ItemName)
	if item.ID == uuid.Nil {
		config.DB.Create(&item)
	} else {
		fmt.Printf("Item %s already exists\n", item.ItemName)
	}

	catalog_item := models.CatalogItem{
		ItemID: item.ID, MarkedPrice: decimal.NewFromFloat(65000.00), DiscountAmount: decimal.NewFromFloat(3000.00),
		SellingPrice: decimal.NewFromFloat(62000.00), BaseModel: baseModel}
	config.DB.Preload("Item").Find(&catalog_item, "item_id = ?", item.ID)
	if catalog_item.ID == uuid.Nil {
		config.DB.Create(&catalog_item)
	} else {
		fmt.Printf("Catalog Item %s already exists\n", catalog_item.Item.ItemName)
	}
	fmt.Println("Successfully loaded the default data")
}
