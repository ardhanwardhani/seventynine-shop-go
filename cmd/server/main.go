package main

import (
	"log"
	"seventynine-shop-go/internal/application"
	"seventynine-shop-go/internal/infrastructure/postgres"
	"seventynine-shop-go/internal/infrastructure/web"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := postgres.NewPostgresDB("postgresql://postgres:12345@localhost:5432/seventynine-db")

	productRepo := postgres.NewPostgresProductRepository(db)
	productService := application.NewProductService(productRepo)

	handler := web.NewProductHandler(productService)
	web.SetupRouter(app, handler)

	log.Fatal(app.Listen(":3000"))
}
