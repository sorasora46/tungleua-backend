package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sorasora46/Tungleua-backend/app/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Get("/find-by-id/:id", handlers.GetUserById)
	user.Get("/exist", handlers.CheckDuplicateUser)
	user.Get("/find-by-email", handlers.GetUserByEmail)
	user.Put("/update/:id", handlers.UpdateUserById)
	user.Post("/", handlers.CreateUser)

	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	// auth.Post("/refresh-token", handlers.RefreshToken)

	store := api.Group("/stores")
	store.Get("/find-by-id/:id", handlers.GetStoreById)
	store.Get("/find-by-user/:id", handlers.GetStoreByUserId)
	store.Get("/exist", handlers.CheckDuplicateStore)
	store.Put("/update/:id", handlers.UpdateStoreById)
	store.Post("/", handlers.CreateStore)
	store.Delete("/delete", handlers.DeleteStoreById)
	store.Get("/populate", handlers.PopulateMap)

	product := api.Group("/products")
	product.Post("/", handlers.CreateProduct)
	product.Get("/find-many-by-id/:id", handlers.GetProducts)
	product.Get("/find-by-id/:id", handlers.GetProductById)
	product.Put("/update/:id", handlers.UpdateProductById)
	product.Delete("/delete/:id", handlers.DeleteProductById)
	product.Get("/images/:id", handlers.GetProductImages)

	order := api.Group("/orders")
	order.Post("/", handlers.CreateOrder)
	order.Get("/find-by-id/:id", handlers.GetOrderById)
	order.Get("/find-by-user/:id", handlers.GetOrders)
	order.Put("/update/:id", handlers.UpdateOrderStatus)
	order.Delete("/delete/:id", handlers.DeleteOrderById)

	discount := api.Group("/discounts")
	discount.Get("/", handlers.GetDiscounts)
	discount.Get("/find-by-id/:id", handlers.GetDiscountById)
}
