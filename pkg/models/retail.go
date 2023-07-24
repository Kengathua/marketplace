package models

import (
	"github.com/Kengathua/marketplace/pkg/common"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Division struct {
	common.Base  `gorm:"embedded"`
	DivisionName string `gorm:"column:division_name" json:"division_name"`
	DivisionCode string `gorm:"column:division_code" json:"division_code"`
	Description  string `gorm:"column:description" json:"description"`
}

type SuperCategory struct {
	common.Base       `gorm:"embedded"`
	DivisionID        string   `gorm:"foreignKey:ID" json:"division_id"`
	Division          Division `gorm:"foreignkey:DivisionID;references:id" json:"division"`
	SuperCategoryName string   `gorm:"column:super_category_name" json:"super_category_name"`
	SuperCategoryCode string   `gorm:"column:super_category_code" json:"super_category_code"`
	Description       string   `gorm:"column:description" json:"description"`
}

type Category struct {
	common.Base     `gorm:"embedded"`
	SuperCategoryID string        `gorm:"foreignKey:ID" json:"super_category_id"`
	SuperCategory   SuperCategory `gorm:"foreignkey:SuperCategoryID;references:id" json:"super_category"`
	CategoryName    string        `gorm:"column:category_name" json:"category_name"`
	CategoryCode    string        `gorm:"column:category_code" json:"category_code"`
	Description     string        `gorm:"column:description" json:"description"`
}

type ItemType struct {
	common.BaseModel `gorm:"embedded"`
	CategoryID       string   `gorm:"foreignKey:ID" json:"category_id"`
	Category         Category `gorm:"foreignkey:CategoryID;references:id" json:"category"`
	TypeName         string   `gorm:"column:type_name" json:"type_name"`
	TypeCode         string   `gorm:"column:type_code" json:"type_code"`
}

type Brand struct {
	common.BaseModel `gorm:"embedded"`
	BrandName        string      `gorm:"column:brand_name" json:"brand_name"`
	BrandCode        string      `gorm:"column:brand_code" json:"brand_code"`
	ItemTypes        []*ItemType `gorm:"many2many:brand_item_types;" json:"item_types"`
}

type BrandItemType struct {
	common.BaseModel `gorm:"embedded"`
	BrandID          string   `gorm:"foreignKey:ID" json:"brand_id"`
	Brand            Brand    `gorm:"foreignkey:BrandID;references:id" json:"brand"`
	ItemTypeID       string   `gorm:"foreignKey:ID" json:"item_type_id"`
	ItemType         ItemType `gorm:"foreignkey:ItemTypeID;references:id" json:"item_type"`
}

type Model struct {
	common.BaseModel `gorm:"embedded"`
	BrandID          string   `gorm:"foreignKey:ID" json:"brand_id"`
	Brand            Brand    `gorm:"foreignkey:BrandID;references:id" json:"brand"`
	ItemTypeID       string   `gorm:"foreignKey:ID" json:"item_type_id"`
	ItemType         ItemType `gorm:"foreignkey:ItemTypeID;references:id" json:"item_type"`
	ModelNumber      string   `gorm:"column:model_number" json:"model_number"`
	ModelCode        string   `gorm:"column:model_code" json:"model_code"`
}

type Item struct {
	common.BaseModel `gorm:"embedded"`
	ModelID          string `gorm:"foreignKey:ID;null" json:"model_id"`
	Model            Model  `gorm:"foreignkey:ModelID;references:id" json:"model"`
	ItemSize         string `gorm:"column:item_size" json:"item_size"`
	ItemName         string `gorm:"column:item_name" json:"item_name"`
	Barcode          string `gorm:"column:barcode" json:"barcode"`
	ItemCode         string `gorm:"column:item_code" json:"item_code"`
	MakeYear         string `gorm:"column:make_year" json:"make_year"`
}

type ItemAttribute struct {
	common.BaseModel `gorm:"embedded"`
	ItemID           string `gorm:"foreignKey:ID" json:"item_id"`
	Item             Item   `gorm:"foreignkey:ItemID;references:id" json:"item"`
	AttributeType    string `gorm:"column:attribute_type" json:"attribute_type"`
	AttributeValue   string `gorm:"column:attribute_value" json:"attribute_value"`
}

type ItemImage struct {
	common.BaseModel `gorm:"embedded"`
	ItemID           string `gorm:"foreignKey:ID" json:"item_id"`
	Item             Item   `gorm:"foreignkey:ItemID;references:id" json:"item"`
	Image            []byte
	IsHeroImage      string `gorm:"column:is_hero_image" json:"is_hero_image"`
}

type CatalogItem struct {
	common.BaseModel `gorm:"embedded"`
	ItemID           string          `gorm:"foreignKey:ID" json:"item_id"`
	Item             Item            `gorm:"foreignkey:ItemID;references:id" json:"item"`
	MarkedPrice      decimal.Decimal `gorm:"type:decimal(10,2);column:marked_price" json:"marked_price"`
	DiscountAmount   decimal.Decimal `gorm:"type:decimal(10,2);column:discount_amount" json:"discount_amount"`
	SellingPrice     decimal.Decimal `gorm:"type:decimal(10,2);column:selling_price" json:"selling_price"`
	ThresholdPrice   decimal.Decimal `gorm:"type:decimal(10,2);column:threshold_price" json:"threshold_price"`
}

// Add owner -> References user
type Cart struct {
	common.BaseModel `gorm:"embedded"`
	CartName         string
	CartCode         string
	OrderGUID        uuid.UUID
	CustomerCartGUID uuid.UUID
}

type CartItem struct {
	common.BaseModel `gorm:"embedded"`
	CartID           uuid.UUID
	Cart             Cart
	CatalogItemID    uuid.UUID
	CatalogItem      CatalogItem
	UniPrice         string
	Quantity         string
	TotalPrice       string
}

type Order struct {
	common.BaseModel  `gorm:"embedded"`
	CartName          string
	CartCode          string
	CartID            uuid.UUID
	Cart              Cart
	CustomerOrderGUID uuid.UUID
}

type OrderItem struct {
	common.BaseModel `gorm:"embedded"`
	OrderID          uuid.UUID
	Order            Order
	CartItemID       uuid.UUID
	CartItem         CartItem
	UniPrice         string
	Quantity         string
	TotalPrice       string
}
