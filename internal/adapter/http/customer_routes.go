package http

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

func SetUpCustomerRoutes(api fiber.Router, customerService port.CustomerService) {
	customerHandler := NewCustomerHandler(customerService)
	customerRoutes := api.Group("/customers")
	customerRoutes.Get("/", customerHandler.GetAllCustomer)
	customerRoutes.Get("/id/:id", customerHandler.GetCustomerByID)
	customerRoutes.Get("/card_uid/:card_uid", customerHandler.GetCustomerByCardUID)
	customerRoutes.Get("/phone/:phone", customerHandler.GetCustomerByPhone)
	customerRoutes.Post("/", customerHandler.CreateCustomer)
	customerRoutes.Put("/:id", customerHandler.UpdateCustomer)
	customerRoutes.Delete("/:id", customerHandler.DeleteCustomer)
}
