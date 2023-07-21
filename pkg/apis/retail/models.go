package retail

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/models"
)

type ModelRequestBody struct {
	BrandID     uuid.UUID `json:"brand_id"`
	ItemTypeID  uuid.UUID `json:"item_type_id"`
	ModelNumber string    `json:"type_name"`
	ModelCode   string    `json:"type_code"`
}

func (h Handler) GetModels(c *fiber.Ctx) error {
	var itemModels []models.Model

	if result := h.DB.Preload("Brand").Preload("ItemType.Category.SuperCategory.Division").Find(&itemModels); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&itemModels)
}

func (h Handler) GetModel(c *fiber.Ctx) error {
	id := c.Params("id")
	var itemModel models.Model

	if result := h.DB.Preload("Brand").Preload("ItemType.Category.SuperCategory.Division").First(&itemModel, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&itemModel)
}

func (h Handler) AddModel(c *fiber.Ctx) error {
	body := ModelRequestBody{}
	// user := c.Locals("user").(models.User)

	// parse body, attach to ModelRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var itemModel models.Model

	itemModel.BrandID = body.BrandID
	itemModel.ItemTypeID = body.ItemTypeID
	itemModel.ModelNumber = body.ModelNumber
	itemModel.ModelCode = body.ModelCode
	// itemModel.CreatedBy = user.ID
	// itemModel.UpdatedBy = user.ID

	if result := h.DB.Create(&itemModel); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.Preload("Brand").Preload("ItemType.Category.SuperCategory.Division").First(&itemModel, "id = ?", itemModel.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&itemModel)
}

func (h Handler) UpdateModel(c *fiber.Ctx) error {
	id := c.Params("id")
	body := ModelRequestBody{}
	// user := c.Locals("user").(models.User)

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var itemModel models.Model

	if result := h.DB.Preload("Brand").Preload("ItemType.Category.SuperCategory.Division").First(&itemModel, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	itemModel.BrandID = body.BrandID
	itemModel.ItemTypeID = body.ItemTypeID
	itemModel.ModelNumber = body.ModelNumber
	itemModel.ModelCode = body.ModelCode
	// itemModel.UpdatedBy = user.ID

	h.DB.Save(&itemModel)

	return c.Status(fiber.StatusOK).JSON(&itemModel)
}

func (h Handler) DeleteModel(c *fiber.Ctx) error {
	id := c.Params("id")
	var itemModel models.Model

	if result := h.DB.First(&itemModel, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&itemModel)

	return c.SendStatus(fiber.StatusOK)
}
