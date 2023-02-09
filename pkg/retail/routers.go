package retail

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matawis/matawis/pkg/retail/controllers"
	"gorm.io/gorm"
)

func RegisterRoutes(url fiber.Router, db *gorm.DB) {
	h := &controllers.Handler{
		DB: db,
	}

	DivisionsRoutes := url.Group("/divisions")
	DivisionsRoutes.Post("/", h.AddDivision)
	DivisionsRoutes.Get("/", h.GetDivisions)
	DivisionsRoutes.Get("/:id", h.GetDivision)
	DivisionsRoutes.Put("/:id", h.UpdateDivision)
	DivisionsRoutes.Delete("/:id", h.DeleteDivision)

	SuperCategoriesRoutes := url.Group("/super_categories")
	SuperCategoriesRoutes.Post("/", h.AddSuperCategory)
	SuperCategoriesRoutes.Get("/", h.GetSuperCategories)
	SuperCategoriesRoutes.Get("/:id", h.GetSuperCategory)
	SuperCategoriesRoutes.Put("/:id", h.UpdateSuperCategory)
	SuperCategoriesRoutes.Delete("/:id", h.DeleteSuperCategory)

	CategoriesRoutes := url.Group("/categories")
	CategoriesRoutes.Post("/", h.AddCategory)
	CategoriesRoutes.Get("/", h.GetCategories)
	CategoriesRoutes.Get("/:id", h.GetCategory)
	CategoriesRoutes.Put("/:id", h.UpdateCategory)
	CategoriesRoutes.Delete("/:id", h.DeleteCategory)
}
