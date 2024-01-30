package main

import (
	"fmt"

	"github.com/aldisypu/go-pos/internal/config"
	"github.com/aldisypu/go-pos/internal/handler"
	"github.com/aldisypu/go-pos/internal/repository"
	"github.com/aldisypu/go-pos/internal/routes"
	"github.com/aldisypu/go-pos/internal/service"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, log)
	app := config.NewFiber(viperConfig)

	productRepo := repository.NewProductRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	productService := service.NewProductService(*productRepo)
	transactionService := service.NewTransactionService(*transactionRepo)

	productHandler := handler.NewProductHandler(*productService)
	transactionHandler := handler.NewTransactionHandler(*transactionService)

	routes.SetupRoutes(app, productHandler, transactionHandler)

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
