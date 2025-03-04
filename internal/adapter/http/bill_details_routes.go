package http

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

func SetUpBill_DetailsRoutes(api fiber.Router, bill_detailsService port.Bill_DetailsService) {
	bill_detailsHandler := NewBill_DetailsHandler(bill_detailsService)
	bill_detailsRoutes := api.Group("/bill_details")
	bill_detailsRoutes.Get("/", bill_detailsHandler.GetAllBill_Details)
	bill_detailsRoutes.Get("/id/:id", bill_detailsHandler.GetBill_DetailsByID)
	bill_detailsRoutes.Post("/", bill_detailsHandler.CreateBill_Details)
	bill_detailsRoutes.Put("/:id", bill_detailsHandler.UpdateBill_Details)
	bill_detailsRoutes.Delete("/:id", bill_detailsHandler.DeleteBill_Details)
}
