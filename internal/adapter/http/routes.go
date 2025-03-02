package http

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, productService port.ProductSevice){
	api := app.Group("/api")

	SetUpProductRoutes(api, productService)
}