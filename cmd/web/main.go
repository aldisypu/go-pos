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

	productService := service.NewProductService(*productRepo)

	productHandler := handler.NewProductHandler(*productService)

	routes.SetupRoutes(app, productHandler)

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
