package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/models"
)

type CategoryRequestBody struct {
	SuperCategoryID uuid.UUID `json:"super_category_id"`
	CategoryName    string    `json:"category_name"`
	CategoryCode    string    `json:"category_code"`
	Description     string    `json:"description"`
}

func (h Handler) GetCategories(c *fiber.Ctx) error {
	var categories []models.Category

	if result := h.DB.Preload("SuperCategory.Division").Find(&categories); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&categories)
}

func (h Handler) GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category

	if result := h.DB.Preload("SuperCategory.Division").First(&category, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&category)
}

func (h Handler) AddCategory(c *fiber.Ctx) error {
	body := CategoryRequestBody{}
	// user := c.Locals("user").(models.User)

	// parse body, attach to CategoryRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var category models.Category

	category.SuperCategoryID = body.SuperCategoryID
	category.CategoryName = body.CategoryName
	category.CategoryCode = body.CategoryCode
	category.Description = body.Description
	// category.CreatedBy = user.ID
	// category.UpdatedBy = user.ID

	if result := h.DB.Create(&category); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.Preload("SuperCategory.Division").First(&category, "id = ?", category.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&category)
}

func (h Handler) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	body := CategoryRequestBody{}
	// user := c.Locals("user").(models.User)

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var category models.Category

	if result := h.DB.Preload("SuperCategory.Division").First(&category, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	category.SuperCategoryID = body.SuperCategoryID
	category.CategoryName = body.CategoryName
	category.CategoryCode = body.CategoryCode
	category.Description = body.Description
	// category.UpdatedBy = user.ID

	h.DB.Save(&category)

	return c.Status(fiber.StatusOK).JSON(&category)
}

func (h Handler) DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category

	if result := h.DB.First(&category, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&category)

	return c.SendStatus(fiber.StatusOK)
}
