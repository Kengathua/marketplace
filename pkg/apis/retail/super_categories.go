package retail

import (
	"github.com/Kengathua/marketplace/pkg/models"
	"github.com/gofiber/fiber/v2"
)

type SuperCategoryRequestBody struct {
	DivisionID        string `json:"division_id"`
	SuperCategoryName string `json:"super_category_name"`
	SuperCategoryCode string `json:"super_category_code"`
	Description       string `json:"description"`
}

func (h Handler) GetSuperCategories(c *fiber.Ctx) error {
	var superCategories []models.SuperCategory

	if result := h.DB.Preload("Division").Find(&superCategories); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&superCategories)
}

func (h Handler) GetSuperCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var superCategory models.SuperCategory

	if result := h.DB.Preload("Division").First(&superCategory, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&superCategory)
}

func (h Handler) AddSuperCategory(c *fiber.Ctx) error {
	body := SuperCategoryRequestBody{}
	// user := c.Locals("user").(models.User)

	// parse body, attach to SuperCategoryRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var superCategory models.SuperCategory

	superCategory.DivisionID = body.DivisionID
	superCategory.SuperCategoryName = body.SuperCategoryName
	superCategory.SuperCategoryCode = body.SuperCategoryCode
	superCategory.Description = body.Description
	// superCategory.CreatedBy = user.ID
	// superCategory.UpdatedBy = user.ID

	// insert new db entry
	if result := h.DB.Create(&superCategory); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.Preload("Division").First(&superCategory, "id = ?", superCategory.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&superCategory)
}

func (h Handler) UpdateSuperCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	body := SuperCategoryRequestBody{}
	// user := c.Locals("user").(models.User)

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var superCategory models.SuperCategory

	if result := h.DB.Preload("Division").First(&superCategory, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	superCategory.DivisionID = body.DivisionID
	superCategory.SuperCategoryName = body.SuperCategoryName
	superCategory.SuperCategoryCode = body.SuperCategoryCode
	superCategory.Description = body.Description
	// superCategory.UpdatedBy = user.ID

	h.DB.Save(&superCategory)

	return c.Status(fiber.StatusOK).JSON(&superCategory)
}

func (h Handler) DeleteSuperCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var superCategory models.SuperCategory

	if result := h.DB.Preload("Division").First(&superCategory, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&superCategory)

	return c.SendStatus(fiber.StatusOK)
}
