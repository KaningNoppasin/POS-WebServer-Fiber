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

type BillHandler struct {
	service port.BillService
}

func NewBillHandler(service port.BillService) *BillHandler {
	return &BillHandler{service: service}
}

func (h *BillHandler) GetAllBill(c *fiber.Ctx) error {
	bills, err := h.service.GetAllBill()
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, bills)
}

func (h *BillHandler) GetBillByID(c *fiber.Ctx) error {
	billId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	bill, err := h.service.GetBillByID(uint(billId))
	if err != nil {
		if errors.Is(err, service.ErrBillNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, bill)
}

func (h *BillHandler) CreateBill(c *fiber.Ctx) error {
	/*
		{
			"customer_id": 1,
			"bill_details": [
				{
					"product_id": 1,
					"quantity": 2
				},
				{
					"product_id": 2,
					"quantity": 1
				}
			]
		}
	*/

	var req entity.CreateBillRequest
	err := c.BodyParser(&req)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = h.service.CreateBill(&req)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
	})

	// var bill entity.Bill
	// // get customer id from form data and convert in to integer
	// customerId, err := strconv.ParseUint(c.FormValue("customer_id"), 10, 64)
	// if err != nil {
	// 	return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	// }

	// // get totalAmount from form data and convert in to integer
	// totalAmount, err := strconv.ParseUint(c.FormValue("total_amount"), 10, 64)
	// if err != nil {
	// 	return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	// }

	// bill = entity.Bill{
	// 	CustomerID:  uint(customerId),
	// 	TotalAmount: uint(totalAmount),
	// }

	// err = h.service.CreateBill(&bill)
	// if err != nil {
	// 	return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	// }
	// return c.Status(fiber.StatusCreated).JSON(fiber.Map{
	// 	"status": "success",
	// })
}

func (h *BillHandler) UpdateBill(c *fiber.Ctx) error {
	var bill entity.Bill
	billId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get customer id from form data and convert in to integer
	customerId, err := strconv.ParseUint(c.FormValue("customer_id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get totalAmount from form data and convert in to integer
	totalAmount, err := strconv.ParseUint(c.FormValue("total_amount"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	bill = entity.Bill{
		CustomerID:  uint(customerId),
		TotalAmount: uint(totalAmount),
	}
	bill.ID = uint(billId)

	err = h.service.UpdateBill(&bill)
	if err != nil {
		if errors.Is(err, service.ErrBillNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *BillHandler) DeleteBill(c *fiber.Ctx) error {
	billId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = h.service.DeleteBill(uint(billId))
	if err != nil {
		if errors.Is(err, service.ErrBillNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
