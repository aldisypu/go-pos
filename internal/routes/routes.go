package routes

import (
	"github.com/aldisypu/go-pos/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, productHandler *handler.ProductHandler, transactionHandler *handler.TransactionHandler) {
	app.Post("/api/products", productHandler.CreateProduct)
	app.Get("/api/products", productHandler.GetAllProducts)
	app.Get("/api/products/:id", productHandler.GetProductByID)
	app.Put("/api/products/:id", productHandler.UpdateProduct)
	app.Delete("/api/products/:id", productHandler.DeleteProduct)

	app.Post("/api/transactions", transactionHandler.CreateTransaction)
	app.Get("/api/transactions", transactionHandler.GetAllTransactions)
	app.Get("/api/transactions/:id", transactionHandler.GetTransactionByID)
	app.Put("/api/transactions/:id", transactionHandler.UpdateTransaction)
	app.Delete("/api/transactions/:id", transactionHandler.DeleteTransaction)
}
