package orders

import (
	"github.com/Kengathua/marketplace/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type CustomerOrderRequestBody struct {
	CustomerCartID string `json:"customer_cart_id"`
	OrderName      string `json:"order_name"`
	OrderCode      string `json:"order_code"`
}

func (h Handler) GetCustomerOrders(c *fiber.Ctx) error {
	var orders []models.CustomerOrder

	if result := h.DB.Find(&orders); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&orders)
}

func (h Handler) GetCustomerOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.CustomerOrder

	if result := h.DB.First(&order, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&order)
}

func (h Handler) AddCustomerOrder(c *fiber.Ctx) error {
	body := CustomerOrderRequestBody{}
	// user := c.Locals("user").(models.User)

	// parse body, attach to CustomerOrderRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var order models.CustomerOrder

	order.CustomerCartID = body.CustomerCartID
	order.OrderName = body.OrderName
	order.OrderCode = body.OrderCode

	if result := h.DB.Create(&order); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.First(&order, "id = ?", order.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&order)
}

func (h Handler) UpdateCustomerOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	body := CustomerOrderRequestBody{}
	// user := c.Locals("user").(models.User)

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var order models.CustomerOrder

	if result := h.DB.First(&order, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	order.CustomerCartID = body.CustomerCartID
	order.OrderName = body.OrderName
	order.OrderCode = body.OrderCode

	h.DB.Save(&order)

	return c.Status(fiber.StatusOK).JSON(&order)
}

func (h Handler) DeleteCustomerOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	var order models.CustomerOrder

	if result := h.DB.First(&order, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&order)

	return c.SendStatus(fiber.StatusOK)
}
