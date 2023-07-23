package orders

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/models"
)

type CustomerCartItemRequestBody struct {
	CustomerCartID string `json:"customer_cart_id"`
	CatalogItemID  string `json:"catalog_item_id"`
	UnitPrice      string `json:"unit_price"`
	Quantity       string `json:"quantity"`
	TotalPrice     string `json:"total_price"`
}

func (h Handler) GetCustomerCartItems(c *fiber.Ctx) error {
	var cartItems []models.CustomerCartItem

	if result := h.DB.Find(&cartItems); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&cartItems)
}

func (h Handler) GetCustomerCartItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var cartItem models.CustomerCartItem

	if result := h.DB.First(&cartItem, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&cartItem)
}

func (h Handler) AddCustomerCartItem(c *fiber.Ctx) error {
	body := CustomerCartItemRequestBody{}
	user := c.Locals("user").(models.User)
	userID, err := uuid.Parse(*user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// parse body, attach to CustomerCartItemRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var cartItem models.CustomerCartItem

	cartItem.CustomerCartID = body.CustomerCartID
	cartItem.CatalogItemID = body.CatalogItemID
	cartItem.UnitPrice = body.UnitPrice
	cartItem.Quantity = body.Quantity
	cartItem.TotalPrice = body.TotalPrice
	cartItem.CreatedBy = userID
	cartItem.UpdatedBy = userID

	if result := h.DB.Create(&cartItem); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.First(&cartItem, "id = ?", cartItem.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&cartItem)
}

func (h Handler) UpdateCustomerCartItem(c *fiber.Ctx) error {
	id := c.Params("id")
	body := CustomerCartItemRequestBody{}
	user := c.Locals("user").(models.User)
	userID, err := uuid.Parse(*user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var cartItem models.CustomerCartItem

	if result := h.DB.First(&cartItem, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	cartItem.CustomerCartID = body.CustomerCartID
	cartItem.CatalogItemID = body.CatalogItemID
	cartItem.UnitPrice = body.UnitPrice
	cartItem.Quantity = body.Quantity
	cartItem.TotalPrice = body.TotalPrice
	cartItem.CreatedBy = userID
	cartItem.UpdatedBy = userID

	h.DB.Save(&cartItem)

	return c.Status(fiber.StatusOK).JSON(&cartItem)
}

func (h Handler) DeleteCustomerCartItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var cartItem models.CustomerCartItem

	if result := h.DB.First(&cartItem, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&cartItem)

	return c.SendStatus(fiber.StatusOK)
}
