package main

import (
	"golang-mongodb-repository/app"
	"golang-mongodb-repository/configs"
	"golang-mongodb-repository/repository"
	"golang-mongodb-repository/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber uygulamasını oluştur
	appRoute := fiber.New()

	// MongoDB veritabanı bağlantısını oluştur
	client := configs.ConnectDatabase()
	productCollection := configs.GetCollection(client, "products")

	// Repository ve Service katmanlarını oluştur
	productRepository := repository.NewProductRepositoryDB(productCollection)
	productService := services.NewProductService(&productRepository)
	productHandler := app.ProductHandler{Service: productService}

	// API rotalarını tanımla
	appRoute.Post("/api/product", productHandler.CreateProduct)
	appRoute.Get("/api/products", productHandler.GetAllProducts)
	appRoute.Delete("/api/product/:id", productHandler.DeleteProduct)

	// Uygulamayı dinle
	appRoute.Listen(":8080")
}
