package http

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

func SetUpBillRoutes(api fiber.Router, billService port.BillService) {
	billHandler := NewBillHandler(billService)
	billRoutes := api.Group("/bills")
	billRoutes.Get("/", billHandler.GetAllBill)
	billRoutes.Get("/id/:id", billHandler.GetBillByID)
	billRoutes.Post("/", billHandler.CreateBill)
	billRoutes.Put("/:id", billHandler.UpdateBill)
	billRoutes.Delete("/:id", billHandler.DeleteBill)
}
