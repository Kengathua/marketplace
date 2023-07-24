package models

import (
	"github.com/Kengathua/marketplace/pkg/common"
)

type CustomerCart struct {
	common.BaseModel
	CustomerID        string
	Customer          Customer
	CartName          string
	CartCode          string
	CustomerOrderGUID *string
}

type CustomerCartItem struct {
	common.BaseModel
	CustomerCartID string
	CustomerCart   CustomerCart
	CatalogItemID  string
	CatalogItem    CatalogItem
	UnitPrice      string
	Quantity       string
	TotalPrice     string
}

type CustomerOrder struct {
	common.BaseModel
	CustomerCartID string
	CustomerCart   CustomerCart
	OrderName      string
	OrderCode      string
}

type CustomerOrderItem struct {
	common.BaseModel
	CustomerOrderID    string
	CustomerOrder      CustomerOrder
	CustomerCartItemID string
	CustomerCartItem   CustomerCartItem
	UnitPrice          string
	Quantity           string
	TotalPrice         string
}
