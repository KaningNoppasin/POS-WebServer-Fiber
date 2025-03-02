package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type StockRepository interface {
	GetAll() ([]entity.Stock, error)
	GetByID(id uint) (*entity.Stock, error)
	Update(stock *entity.Stock) error
}
