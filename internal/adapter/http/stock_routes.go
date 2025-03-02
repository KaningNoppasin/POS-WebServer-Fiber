package http

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

func SetUpStockRoutes(api fiber.Router, stockService port.StockService){
	stockHandler := NewStockHandler(stockService)
	stockRoutes := api.Group("/Stocks")
	stockRoutes.Get("/", stockHandler.GetAllStock)
	stockRoutes.Get("/id/:id", stockHandler.GetStockByID)
	stockRoutes.Put("/:id", stockHandler.UpdateStock)
}