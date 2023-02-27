package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/models"
)

type BrandItemTypeRequestBody struct {
	BrandID    uuid.UUID `json:"brand_id"`
	ItemTypeID uuid.UUID `json:"item_type_id"`
}

func (h Handler) GetBrandItemTypes(c *fiber.Ctx) error {
	var brandItemTypes []models.BrandItemType

	if result := h.DB.Preload("Brand").Preload("ItemType.Category.SuperCategory.Division").Find(&brandItemTypes); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&brandItemTypes)
}

func (h Handler) GetBrandItemType(c *fiber.Ctx) error {
	id := c.Params("id")
	var brandItemType models.BrandItemType

	if result := h.DB.Preload("ItemType.Category.SuperCategory.Division").First(&brandItemType, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&brandItemType)
}

func (h Handler) AddBrandItemType(c *fiber.Ctx) error {
	body := BrandItemTypeRequestBody{}
	// user := c.Locals("user").(models.User)

	// parse body, attach to BrandItemTypeRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var brandItemType models.BrandItemType

	brandItemType.BrandID = body.BrandID
	brandItemType.ItemTypeID = body.ItemTypeID
	// brandItemType.CreatedBy = user.ID
	// brandItemType.UpdatedBy = user.ID

	if result := h.DB.Create(&brandItemType); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.Preload("ItemType.Category.SuperCategory.Division").First(&brandItemType, "id = ?", brandItemType.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&brandItemType)
}

func (h Handler) UpdateBrandItemType(c *fiber.Ctx) error {
	id := c.Params("id")
	body := BrandItemTypeRequestBody{}
	// user := c.Locals("user").(models.User)

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var brandItemType models.BrandItemType

	if result := h.DB.Preload("ItemType.Category.SuperCategory.Division").First(&brandItemType, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	brandItemType.BrandID = body.BrandID
	brandItemType.ItemTypeID = body.ItemTypeID
	// brandItemType.UpdatedBy = user.ID

	h.DB.Save(&brandItemType)

	return c.Status(fiber.StatusOK).JSON(&brandItemType)
}

func (h Handler) DeleteBrandItemType(c *fiber.Ctx) error {
	id := c.Params("id")
	var brandItemType models.BrandItemType

	if result := h.DB.First(&brandItemType, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&brandItemType)

	return c.SendStatus(fiber.StatusOK)
}
