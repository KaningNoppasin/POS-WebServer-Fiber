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

type Bill_DetailsHandler struct {
	service port.Bill_DetailsService
}

func NewBill_DetailsHandler(service port.Bill_DetailsService) *Bill_DetailsHandler {
	return &Bill_DetailsHandler{service: service}
}

func (h *Bill_DetailsHandler) GetAllBill_Details(c *fiber.Ctx) error {
	bill_details, err := h.service.GetAllBill_Details()
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, bill_details)
}

func (h *Bill_DetailsHandler) GetBill_DetailsByID(c *fiber.Ctx) error {
	bill_detailsId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	bill_detail, err := h.service.GetBill_DetailsByID(uint(bill_detailsId))
	if err != nil {
		if errors.Is(err, service.ErrBill_DetailsNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, bill_detail)
}

func (h *Bill_DetailsHandler) CreateBill_Details(c *fiber.Ctx) error {
	var bill_details entity.Bill_Details
	// get bill id from form data and convert in to integer
	billId, err := strconv.ParseUint(c.FormValue("bill_id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get product id from form data and convert in to integer
	productId, err := strconv.ParseUint(c.FormValue("product_id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get quantity from form data and convert in to integer
	quantity, err := strconv.ParseUint(c.FormValue("quantity"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get total from form data and convert in to integer
	total, err := strconv.ParseUint(c.FormValue("total"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	bill_details = entity.Bill_Details{
		BillID:  uint(billId),
		ProductID:  uint(productId),
		Quantity: uint(quantity),
		Total: uint(total),
	}

	err = h.service.CreateBill_Details(&bill_details)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *Bill_DetailsHandler) UpdateBill_Details(c *fiber.Ctx) error {
	var bill_details entity.Bill_Details
	bill_detailsId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get bill id from form data and convert in to integer
	billId, err := strconv.ParseUint(c.FormValue("bill_id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get product id from form data and convert in to integer
	productId, err := strconv.ParseUint(c.FormValue("product_id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get quantity from form data and convert in to integer
	quantity, err := strconv.ParseUint(c.FormValue("quantity"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	// get total from form data and convert in to integer
	total, err := strconv.ParseUint(c.FormValue("total"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	bill_details = entity.Bill_Details{
		BillID:  uint(billId),
		ProductID:  uint(productId),
		Quantity: uint(quantity),
		Total: uint(total),
	}
	bill_details.ID = uint(bill_detailsId)

	err = h.service.UpdateBill_Details(&bill_details)
	if err != nil {
		if errors.Is(err, service.ErrBill_DetailsNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}

func (h *Bill_DetailsHandler) DeleteBill_Details(c *fiber.Ctx) error {
	bill_detailsId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	err = h.service.DeleteBill_Details(uint(bill_detailsId))
	if err != nil {
		if errors.Is(err, service.ErrBill_DetailsNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}
