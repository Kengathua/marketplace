package retail

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matawis/matawis/pkg/models"
)

type ItemRequestBody struct {
	ItemTypeID uuid.UUID `json:"item_type_id"`
	BrandID    uuid.UUID `json:"brand_id"`
	ModelID    uuid.UUID `json:"model_id"`
	ItemName   string    `json:"item_name"`
	ItemCode   string    `json:"item_code"`
	ItemSize   string    `json:"item_size"`
	Barcode    string    `json:"barcode"`
	MakeYear   string    `json:"make_year"`
}

func (h Handler) GetItems(c *fiber.Ctx) error {
	var items []models.Item

	if result := h.DB.Preload("ItemType.Category.SuperCategory.Division").Preload("Brand").Preload("Model.Brand").Preload("Model.ItemType.Category.SuperCategory.Division").Find(&items); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&items)
}

func (h Handler) GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item

	if result := h.DB.Preload("ItemType.Category.SuperCategory.Division").Preload("Brand").Preload("Model.Brand").Preload("Model.ItemType.Category.SuperCategory.Division").First(&item, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&item)
}

func (h Handler) AddItem(c *fiber.Ctx) error {
	body := ItemRequestBody{}
	// user := c.Locals("user").(models.User)

	// parse body, attach to ItemRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var item models.Item

	item.ItemTypeID = body.ItemTypeID
	item.BrandID = body.BrandID
	item.ModelID = body.ModelID
	item.ItemName = body.ItemName
	item.ItemCode = body.ItemCode
	item.ItemSize = body.ItemSize
	item.Barcode = body.Barcode
	item.MakeYear = body.MakeYear
	// item.CreatedBy = user.ID
	// item.UpdatedBy = user.ID

	if result := h.DB.Create(&item); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	if result := h.DB.Preload("ItemType.Category.SuperCategory.Division").Preload("Brand").Preload("Model.Brand").Preload("Model.ItemType.Category.SuperCategory.Division").First(&item, "id = ?", item.ID); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&item)
}

func (h Handler) UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	body := ItemRequestBody{}
	// user := c.Locals("user").(models.User)

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var item models.Item

	if result := h.DB.Preload("ItemType.Category.SuperCategory.Division").Preload("Brand").Preload("Model.Brand").Preload("Model.ItemType.Category.SuperCategory.Division").First(&item, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	item.ItemTypeID = body.ItemTypeID
	item.BrandID = body.BrandID
	item.ModelID = body.ModelID
	item.ItemName = body.ItemName
	item.ItemCode = body.ItemCode
	item.ItemSize = body.ItemSize
	item.Barcode = body.Barcode
	item.MakeYear = body.MakeYear
	// item.UpdatedBy = user.ID

	h.DB.Save(&item)

	return c.Status(fiber.StatusOK).JSON(&item)
}

func (h Handler) DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item

	if result := h.DB.First(&item, "id = ?", id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	h.DB.Delete(&item)

	return c.SendStatus(fiber.StatusOK)
}
