package apis

import (
	"github.com/Kengathua/marketplace/pkg/apis/orders"
	"github.com/Kengathua/marketplace/pkg/apis/retail"
	"github.com/gofiber/fiber/v2"

	// "github.com/Kengathua/marketplace/pkg/apis/users"
	"gorm.io/gorm"
)

func RegisterAPIRoutes(url fiber.Router, db *gorm.DB) {
	v1 := url.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})
	ordersURL := v1.Group("/orders", func(c *fiber.Ctx) error { // middleware for /api/v1/orders
		c.Set("Version", "v1")
		return c.Next()
	})
	retailURL := v1.Group("/retail", func(c *fiber.Ctx) error { // middleware for /api/v1/retail
		c.Set("Version", "v1")
		return c.Next()
	})

	orders.RegisterOrderRoutes(ordersURL, db)  // /api/v1/orders
	retail.RegisterRetailRoutes(retailURL, db) // /api/v1/retail
}
