package retail

import (
	"github.com/Kengathua/marketplace/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BrandRequestBody struct {
	BrandName string `json:"brand_name"`
	BrandCode string `json:"brand_code"`
}

func (h Handler) GetBrands(c *fiber.Ctx) error {
	var brands []models.Brand

	if result := h.DB.Find(&brands); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&brands)
}

func (h Handler) GetBrand(c *fiber.Ctx) error {
	id := c.Params("id")
	var brand models.Brand

	if result := h.DB.First(&brand, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&brand)
}

func (h Handler) AddBrand(c *fiber.Ctx) error {
	body := BrandRequestBody{}
	user := c.Locals("user").(models.User)
	userID, err := uuid.Parse(*user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// parse body, attach to BrandRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var brand models.Brand

	brand.BrandName = body.BrandName
	brand.BrandCode = body.BrandCode
	brand.CreatedBy = userID
	brand.UpdatedBy = userID

	if result := h.DB.Create(&brand); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.First(&brand, "id = ?", brand.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&brand)
}

func (h Handler) UpdateBrand(c *fiber.Ctx) error {
	id := c.Params("id")
	body := BrandRequestBody{}
	user := c.Locals("user").(models.User)
	userID, err := uuid.Parse(*user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var brand models.Brand

	if result := h.DB.First(&brand, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	brand.BrandName = body.BrandName
	brand.BrandCode = body.BrandCode
	brand.UpdatedBy = userID

	h.DB.Save(&brand)

	return c.Status(fiber.StatusOK).JSON(&brand)
}

func (h Handler) DeleteBrand(c *fiber.Ctx) error {
	id := c.Params("id")
	var brand models.Brand

	if result := h.DB.First(&brand, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&brand)

	return c.SendStatus(fiber.StatusOK)
}
