package app

import (
	"golang-mongodb-repository/models"
	"golang-mongodb-repository/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHandler struct {
	Service services.ProductService
}

func (handler *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := handler.Service.InsertProduct(product)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (handler *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	result, err := handler.Service.GetAllProducts()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (handler *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, err := primitive.ObjectIDFromHex(query)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	result, err := handler.Service.DeleteProduct(cnv)

	if err != nil || !result {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false, "error": "Product not found or could not be deleted"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}
