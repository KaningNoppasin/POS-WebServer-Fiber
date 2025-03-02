package service

import (
	"errors"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"gorm.io/gorm"
)

type StockService struct {
	repo port.StockRepository
}

func NewStockService(repo port.StockRepository) port.StockService {
	return &StockService{repo: repo}
}

func (s *StockService) GetAllStock() ([]entity.Stock, error) {
	stocks, err := s.repo.GetAll()
	if err != nil {
		return nil, ErrFailedToRetrieveStock
	}
	return stocks, nil
}

func (s *StockService) GetStockByID(id uint) (*entity.Stock, error) {
	stock, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStockNotFound
		}
		return nil, ErrFailedToRetrieveStock
	}
	return stock, nil
}

func (s *StockService) UpdateStock(stock *entity.Stock) error {
	_, err := s.repo.GetByID(stock.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrStockNotFound
		}
		return ErrFailedToRetrieveStock
	}

	err = s.repo.Update(stock)
	if err != nil {
		return ErrFailedToCreateStock
	}
	return nil
}
