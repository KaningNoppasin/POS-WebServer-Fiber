package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type StockService interface {
	GetAllStock() ([]entity.Stock, error)
	GetStockByID(id uint) (*entity.Stock, error)
	UpdateStock(stock *entity.Stock) error
}
