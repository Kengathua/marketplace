package orders

import (
	"github.com/Kengathua/marketplace/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type CustomerOrderItemRequestBody struct {
	CustomerOrderID    string `json:"customer_order_id"`
	CustomerCartItemID string `json:"customer_cart_id"`
	UnitPrice          string `json:"unit_price"`
	Quantity           string `json:"quantity_id"`
	TotalPrice         string `json:"total_price"`
}

func (h Handler) GetCustomerOrderItems(c *fiber.Ctx) error {
	var orderItems []models.CustomerOrderItem

	if result := h.DB.Find(&orderItems); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&orderItems)
}

func (h Handler) GetCustomerOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var orderItem models.CustomerOrderItem

	if result := h.DB.First(&orderItem, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&orderItem)
}

func (h Handler) AddCustomerOrderItem(c *fiber.Ctx) error {
	body := CustomerOrderItemRequestBody{}
	// user := c.Locals("user").(models.User)

	// parse body, attach to CustomerOrderItemRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var orderItem models.CustomerOrderItem

	orderItem.CustomerOrderID = body.CustomerOrderID
	orderItem.CustomerCartItemID = body.CustomerCartItemID
	orderItem.UnitPrice = body.UnitPrice
	orderItem.Quantity = body.Quantity
	orderItem.TotalPrice = body.TotalPrice

	if result := h.DB.Create(&orderItem); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.First(&orderItem, "id = ?", orderItem.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&orderItem)
}

func (h Handler) UpdateCustomerOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")
	body := CustomerOrderItemRequestBody{}
	// user := c.Locals("user").(models.User)

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var orderItem models.CustomerOrderItem

	if result := h.DB.First(&orderItem, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	orderItem.CustomerOrderID = body.CustomerOrderID
	orderItem.CustomerCartItemID = body.CustomerCartItemID
	orderItem.UnitPrice = body.UnitPrice
	orderItem.Quantity = body.Quantity
	orderItem.TotalPrice = body.TotalPrice

	h.DB.Save(&orderItem)

	return c.Status(fiber.StatusOK).JSON(&orderItem)
}

func (h Handler) DeleteCustomerOrderItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var orderItem models.CustomerOrderItem

	if result := h.DB.First(&orderItem, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&orderItem)

	return c.SendStatus(fiber.StatusOK)
}
