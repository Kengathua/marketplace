package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/matawis/matawis/pkg/config"
	"github.com/matawis/matawis/pkg/retail"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	config.ConnectToDb()
	api := app.Group("/api")

	v1 := api.Group("/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})
	retailURL := v1.Group("retail", func(c *fiber.Ctx) error { // middleware for /api/v1/retail
		c.Set("Version", "v1")
		return c.Next()
	})
	v1.Group("customers", func(c *fiber.Ctx) error { // middleware for /api/v1/customers
		c.Set("Version", "v1")
		return c.Next()
	})

	retail.RegisterRoutes(retailURL, config.DB)
	app.Listen(":3000")
}
