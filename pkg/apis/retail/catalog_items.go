package retail

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/models"
	"github.com/shopspring/decimal"
)

type CatalogItemRequestBody struct {
	ItemID         uuid.UUID       `json:"item_id"`
	MarkedPrice    decimal.Decimal `json:"marked_price"`
	DiscountAmount decimal.Decimal `json:"discount_amount"`
	SellingPrice   decimal.Decimal `json:"selling_price"`
	ThresholdPrice decimal.Decimal `json:"threshold_price"`
}

func (h Handler) GetCatalogItems(c *fiber.Ctx) error {
	var catalogItems []models.CatalogItem

	if result := h.DB.Preload("Item.Brand").Preload("Item.Model.Brand").Find(&catalogItems); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&catalogItems)
}

func (h Handler) GetCatalogItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var catalogItem models.CatalogItem

	if result := h.DB.Preload("Item.Brand").Preload("Item.Model.Brand").First(&catalogItem, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&catalogItem)
}

func (h Handler) AddCatalogItem(c *fiber.Ctx) error {
	body := CatalogItemRequestBody{}
	// user := c.Locals("user").(models.User)

	// parse body, attach to CatalogItemRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var catalogItem models.CatalogItem

	catalogItem.ItemID = body.ItemID
	catalogItem.MarkedPrice = body.MarkedPrice
	catalogItem.DiscountAmount = body.DiscountAmount
	catalogItem.SellingPrice = body.SellingPrice
	catalogItem.ThresholdPrice = body.ThresholdPrice
	// catalogItem.CreatedBy = user.ID
	// catalogItem.UpdatedBy = user.ID

	if result := h.DB.Create(&catalogItem); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.Preload("Item.Brand").Preload("Item.Model.Brand").First(&catalogItem, "id = ?", catalogItem.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&catalogItem)
}

func (h Handler) UpdateCatalogItem(c *fiber.Ctx) error {
	id := c.Params("id")
	body := CatalogItemRequestBody{}
	// user := c.Locals("user").(models.User)

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var catalogItem models.CatalogItem

	if result := h.DB.Preload("Item.Brand").Preload("Item.Model.Brand").First(&catalogItem, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	catalogItem.ItemID = body.ItemID
	catalogItem.MarkedPrice = body.MarkedPrice
	catalogItem.DiscountAmount = body.DiscountAmount
	catalogItem.SellingPrice = body.SellingPrice
	catalogItem.ThresholdPrice = body.ThresholdPrice
	// catalogItem.UpdatedBy = user.ID

	h.DB.Save(&catalogItem)

	return c.Status(fiber.StatusOK).JSON(&catalogItem)
}

func (h Handler) DeleteCatalogItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var catalogItem models.CatalogItem

	if result := h.DB.First(&catalogItem, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&catalogItem)

	return c.SendStatus(fiber.StatusOK)
}
