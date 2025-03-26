package http

import (
	"errors"
	"strconv"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/service"
	"github.com/KaningNoppasin/Web-Server-Fiber/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	service port.CustomerService
}

func NewCustomerHandler(service port.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) GetAllCustomer(c *fiber.Ctx) error {
	customers, err := h.service.GetAllCustomer()
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, customers)
}

func (h *CustomerHandler) GetCustomerByID(c *fiber.Ctx) error {
	customerId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	customer, err := h.service.GetCustomerByID(uint(customerId))
	if err != nil {
		if errors.Is(err, service.ErrCustomerNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, customer)
}

func (h *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	customer := entity.Customer{
		CustomerName: c.FormValue("customer_name"),
		Phone:        c.FormValue("phone"),
		Email:        c.FormValue("email"),
		CardUID:      c.FormValue("cardUID"),
	}

	err := h.service.CreateCustomer(&customer)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *CustomerHandler) UpdateCustomer(c *fiber.Ctx) error {
	var customer entity.Customer
	customerId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	customer = entity.Customer{
		CustomerName: c.FormValue("customer_name"),
		Phone:        c.FormValue("phone"),
		Email:        c.FormValue("email"),
		CardUID:      c.FormValue("cardUID"),
	}
	customer.ID = uint(customerId)

	err = h.service.UpdateCustomer(&customer)
	if err != nil {
		if errors.Is(err, service.ErrCustomerNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *CustomerHandler) DeleteCustomer(c *fiber.Ctx) error {
	customerId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = h.service.DeleteCustomer(uint(customerId))
	if err != nil {
		if errors.Is(err, service.ErrCustomerNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
