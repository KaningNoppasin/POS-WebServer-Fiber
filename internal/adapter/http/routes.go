package http

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(
	app *fiber.App,
	productService port.ProductService,
	stockService port.StockService,
	customerService port.CustomerService,
	billService port.BillService,
) {
	api := app.Group("/api")

	SetUpProductRoutes(api, productService)
	SetUpStockRoutes(api, stockService)
	SetUpCustomerRoutes(api, customerService)
	SetUpBillRoutes(api, billService)
}
