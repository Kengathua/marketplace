package orders

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterOrderRoutes(url fiber.Router, db *gorm.DB) {
	h := &Handler{
		DB: db,
	}

	CustomerCartsRoutes := url.Group("/customer_carts")
	CustomerCartsRoutes.Post("/", h.AddCustomerCart)
	CustomerCartsRoutes.Get("/", h.GetCustomerCarts)
	CustomerCartsRoutes.Get("/:id", h.GetCustomerCart)
	CustomerCartsRoutes.Put("/:id", h.UpdateCustomerCart)
	CustomerCartsRoutes.Delete("/:id", h.DeleteCustomerCart)

	CustomerCartItemsRoutes := url.Group("/customer_cart_items")
	CustomerCartItemsRoutes.Post("/", h.AddCustomerCartItem)
	CustomerCartItemsRoutes.Get("/", h.GetCustomerCartItems)
	CustomerCartItemsRoutes.Get("/:id", h.GetCustomerCartItem)
	CustomerCartItemsRoutes.Put("/:id", h.UpdateCustomerCartItem)
	CustomerCartItemsRoutes.Delete("/:id", h.DeleteCustomerCartItem)

	CustomerOrdersRoutes := url.Group("/customer_orders")
	CustomerOrdersRoutes.Post("/", h.AddCustomerOrder)
	CustomerOrdersRoutes.Get("/", h.GetCustomerOrders)
	CustomerOrdersRoutes.Get("/:id", h.GetCustomerOrder)
	CustomerOrdersRoutes.Put("/:id", h.UpdateCustomerOrder)
	CustomerOrdersRoutes.Delete("/:id", h.DeleteCustomerOrder)

	CustomerOrderItemsRoutes := url.Group("/customer_order_items")
	CustomerOrderItemsRoutes.Post("/", h.AddCustomerOrderItem)
	CustomerOrderItemsRoutes.Get("/", h.GetCustomerOrderItems)
	CustomerOrderItemsRoutes.Get("/:id", h.GetCustomerOrderItem)
	CustomerOrderItemsRoutes.Put("/:id", h.UpdateCustomerOrderItem)
	CustomerOrderItemsRoutes.Delete("/:id", h.DeleteCustomerOrderItem)
}
