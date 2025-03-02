package http

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

func SetUpProductRoutes(api fiber.Router, productService port.ProductSevice){
	productHandler := NewProductHandler(productService)
	productRoutes := api.Group("/products")
	productRoutes.Get("/", productHandler.GetAllProduct)
	productRoutes.Get("/id/:id", productHandler.GetProductByID)
	productRoutes.Get("/barcode/:barcode", productHandler.GetProductByBarcode)
	productRoutes.Post("/", productHandler.CreateProduct)
	productRoutes.Put("/:id", productHandler.UpdateProduct)
	productRoutes.Delete("/:id", productHandler.DeleteProduct)
}