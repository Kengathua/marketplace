package orders

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/models"
)

type CustomerCartRequestBody struct {
	CustomerID string `json:"customer_id"`
	CartName   string `json:"cart_name"`
	CartCode   string `json:"cart_code"`
}

func (h Handler) GetCustomerCarts(c *fiber.Ctx) error {
	var carts []models.CustomerCart

	if result := h.DB.Find(&carts); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&carts)
}

func (h Handler) GetCustomerCart(c *fiber.Ctx) error {
	id := c.Params("id")
	var cart models.CustomerCart

	if result := h.DB.First(&cart, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&cart)
}

func (h Handler) AddCustomerCart(c *fiber.Ctx) error {
	body := CustomerCartRequestBody{}
	user := c.Locals("user").(models.User)
	userID, err := uuid.Parse(*user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// parse body, attach to CustomerCartRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var cart models.CustomerCart

	cart.CustomerID = body.CustomerID
	cart.CartName = body.CartName
	cart.CartCode = body.CartCode
	cart.CreatedBy = userID
	cart.UpdatedBy = userID

	if result := h.DB.Create(&cart); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.First(&cart, "id = ?", cart.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&cart)
}

func (h Handler) UpdateCustomerCart(c *fiber.Ctx) error {
	id := c.Params("id")
	body := CustomerCartRequestBody{}
	user := c.Locals("user").(models.User)
	userID, err := uuid.Parse(*user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var cart models.CustomerCart

	if result := h.DB.First(&cart, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	cart.CustomerID = body.CustomerID
	cart.CartName = body.CartName
	cart.CartCode = body.CartCode
	cart.UpdatedBy = userID

	h.DB.Save(&cart)

	return c.Status(fiber.StatusOK).JSON(&cart)
}

func (h Handler) DeleteCustomerCart(c *fiber.Ctx) error {
	id := c.Params("id")
	var cart models.CustomerCart

	if result := h.DB.First(&cart, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&cart)

	return c.SendStatus(fiber.StatusOK)
}
