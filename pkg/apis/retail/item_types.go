package retail

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/models"
)

type ItemTypeRequestBody struct {
	CategoryID uuid.UUID `json:"category_id"`
	TypeName   string    `json:"type_name"`
	TypeCode   string    `json:"type_code"`
}

func (h Handler) GetItemTypes(c *fiber.Ctx) error {
	var itemTypes []models.ItemType

	if result := h.DB.Preload("Category.SuperCategory.Division").Find(&itemTypes); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&itemTypes)
}

func (h Handler) GetItemType(c *fiber.Ctx) error {
	id := c.Params("id")
	var itemType models.ItemType

	if result := h.DB.Preload("Category.SuperCategory.Division").First(&itemType, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&itemType)
}

func (h Handler) AddItemType(c *fiber.Ctx) error {
	body := ItemTypeRequestBody{}
	// user := c.Locals("user").(models.User)

	// parse body, attach to ItemTypeRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var itemType models.ItemType

	itemType.CategoryID = body.CategoryID
	itemType.TypeName = body.TypeName
	itemType.TypeCode = body.TypeCode
	// itemType.CreatedBy = user.ID
	// itemType.UpdatedBy = user.ID

	if result := h.DB.Create(&itemType); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.Preload("Category.SuperCategory.Division").First(&itemType, "id = ?", itemType.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&itemType)
}

func (h Handler) UpdateItemType(c *fiber.Ctx) error {
	id := c.Params("id")
	body := ItemTypeRequestBody{}
	// user := c.Locals("user").(models.User)

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var itemType models.ItemType

	if result := h.DB.Preload("Category.SuperCategory.Division").First(&itemType, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	itemType.CategoryID = body.CategoryID
	itemType.TypeName = body.TypeName
	itemType.TypeCode = body.TypeCode
	// itemType.UpdatedBy = user.ID

	h.DB.Save(&itemType)

	return c.Status(fiber.StatusOK).JSON(&itemType)
}

func (h Handler) DeleteItemType(c *fiber.Ctx) error {
	id := c.Params("id")
	var itemType models.ItemType

	if result := h.DB.First(&itemType, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&itemType)

	return c.SendStatus(fiber.StatusOK)
}
