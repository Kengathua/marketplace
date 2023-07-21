package retail

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRetailRoutes(url fiber.Router, db *gorm.DB) {
	h := &Handler{
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

	ItemTypesRoutes := url.Group("/item_types")
	ItemTypesRoutes.Post("/", h.AddItemType)
	ItemTypesRoutes.Get("/", h.GetItemTypes)
	ItemTypesRoutes.Get("/:id", h.GetItemType)
	ItemTypesRoutes.Put("/:id", h.UpdateItemType)
	ItemTypesRoutes.Delete("/:id", h.DeleteItemType)

	BrandsRoutes := url.Group("/brands")
	BrandsRoutes.Post("/", h.AddBrand)
	BrandsRoutes.Get("/", h.GetBrands)
	BrandsRoutes.Get("/:id", h.GetBrand)
	BrandsRoutes.Put("/:id", h.UpdateBrand)
	BrandsRoutes.Delete("/:id", h.DeleteBrand)

	BrandItemTypesRoutes := url.Group("/brand_item_types")
	BrandItemTypesRoutes.Post("/", h.AddBrandItemType)
	BrandItemTypesRoutes.Get("/", h.GetBrandItemTypes)
	BrandItemTypesRoutes.Get("/:id", h.GetBrandItemType)
	BrandItemTypesRoutes.Put("/:id", h.UpdateBrandItemType)
	BrandItemTypesRoutes.Delete("/:id", h.DeleteBrandItemType)

	ModelsRoutes := url.Group("/models")
	ModelsRoutes.Post("/", h.AddModel)
	ModelsRoutes.Get("/", h.GetModels)
	ModelsRoutes.Get("/:id", h.GetModel)
	ModelsRoutes.Put("/:id", h.UpdateModel)
	ModelsRoutes.Delete("/:id", h.DeleteModel)

	ItemsRoutes := url.Group("/items")
	ItemsRoutes.Post("/", h.AddItem)
	ItemsRoutes.Get("/", h.GetItems)
	ItemsRoutes.Get("/:id", h.GetItem)
	ItemsRoutes.Put("/:id", h.UpdateItem)
	ItemsRoutes.Delete("/:id", h.DeleteItem)

	CatalogItemsRoutes := url.Group("/catalog_items")
	CatalogItemsRoutes.Post("/", h.AddCatalogItem)
	CatalogItemsRoutes.Get("/", h.GetCatalogItems)
	CatalogItemsRoutes.Get("/:id", h.GetCatalogItem)
	CatalogItemsRoutes.Put("/:id", h.UpdateCatalogItem)
	CatalogItemsRoutes.Delete("/:id", h.DeleteCatalogItem)
}
