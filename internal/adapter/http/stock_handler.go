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

type StockHandler struct {
	service port.StockService
}

func NewStockHandler(service port.StockService) *StockHandler {
	return &StockHandler{service: service}
}

func (h *StockHandler) GetAllStock(c *fiber.Ctx) error {
	stocks, err := h.service.GetAllStock()
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, stocks)
}

func (h *StockHandler) GetStockByID(c *fiber.Ctx) error {
	stockId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	stock, err := h.service.GetStockByID(uint(stockId))
	if err != nil {
		if errors.Is(err, service.ErrStockNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, stock)
}

func (h *StockHandler) UpdateStock(c *fiber.Ctx) error {
	var stock entity.Stock
	stockId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	stockQuantity, err := strconv.ParseUint(c.FormValue("quantity"), 10, 64)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	stock.ID = uint(stockId)
	stock.Quantity = uint(stockQuantity)

	err = h.service.UpdateStock(&stock)
	if err != nil {
		if errors.Is(err, service.ErrStockNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}