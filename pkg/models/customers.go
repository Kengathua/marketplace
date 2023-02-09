package models

import (
	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/common"
)

type Customer struct {
	common.BioData
	CustomerNumber string
}

type CustomerCart struct {
	common.BaseModel
	CustomerID        uuid.UUID
	Customer          Customer
	CartName          string
	CartCode          string
	CustomerOrderGUID uuid.UUID
}

type CustomerCartItem struct {
	common.BaseModel
	CustomerCartID uuid.UUID
	CustomerCart   CustomerCart
	CatalogItemID  uuid.UUID
	CatalogItem    CatalogItem
	UniPrice       string
	Quantity       string
	TotalPrice     string
}

type CustomerOrder struct {
	common.BaseModel
	CustomerCartID uuid.UUID
	CustomerCart   CustomerCart
	CustomerID     uuid.UUID
	Customer       Customer
	OrderName      string
	OrderCode      string
}

type CustomerOrderItem struct {
	common.BaseModel
	CustomerOrderID    uuid.UUID
	CustomerOrder      CustomerOrder
	CustomerCartItemID uuid.UUID
	CustomerCartItem   CustomerCartItem
	UniPrice           string
	Quantity           string
	TotalPrice         string
}
