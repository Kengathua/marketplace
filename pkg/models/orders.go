package models

import (
	"github.com/matawis/matawis/pkg/common"
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
	CustomerID     string
	Customer       Customer
	OrderName      string
	OrderCode      string
}

type CustomerOrderItem struct {
	common.BaseModel
	CustomerOrderID    string
	CustomerOrder      CustomerOrder
	CustomerCartItemID string
	CustomerCartItem   CustomerCartItem
	UniPrice           string
	Quantity           string
	TotalPrice         string
}
