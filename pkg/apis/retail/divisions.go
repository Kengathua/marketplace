package retail

import (
	"github.com/Kengathua/marketplace/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DivisionRequestBody struct {
	DivisionName string `json:"division_name"`
	DivisionCode string `json:"division_code"`
	Description  string `json:"description"`
}

func (h Handler) GetDivisions(c *fiber.Ctx) error {
	var divisions []models.Division

	if result := h.DB.Find(&divisions); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&divisions)
}

func (h Handler) GetDivision(c *fiber.Ctx) error {
	id := c.Params("id")
	var division models.Division

	if result := h.DB.First(&division, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&division)
}

func (h Handler) AddDivision(c *fiber.Ctx) error {
	body := DivisionRequestBody{}
	user := c.Locals("user").(models.User)
	userID, err := uuid.Parse(*user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// parse body, attach to DivisionRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var division models.Division

	division.DivisionName = body.DivisionName
	division.DivisionCode = body.DivisionCode
	division.Description = body.Description
	division.CreatedBy = userID
	division.UpdatedBy = userID

	// insert new db entry
	if result := h.DB.Create(&division); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&division)
}

func (h Handler) UpdateDivision(c *fiber.Ctx) error {
	id := c.Params("id")
	body := DivisionRequestBody{}
	user := c.Locals("user").(models.User)
	userID, err := uuid.Parse(*user.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var division models.Division

	if result := h.DB.First(&division, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	division.DivisionName = body.DivisionName
	division.DivisionCode = body.DivisionCode
	division.Description = body.Description
	division.UpdatedBy = userID

	// save division
	h.DB.Save(&division)

	return c.Status(fiber.StatusOK).JSON(&division)
}

func (h Handler) DeleteDivision(c *fiber.Ctx) error {
	id := c.Params("id")
	var division models.Division

	if result := h.DB.First(&division, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// delete division from db
	h.DB.Delete(&division)

	return c.SendStatus(fiber.StatusOK)
}
