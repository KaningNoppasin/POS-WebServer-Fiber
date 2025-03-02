package db

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"gorm.io/gorm"
)

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) port.StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) GetAll() ([]entity.Stock, error) {
	var stocks []entity.Stock
	err := r.db.Find(&stocks).Error
	return stocks, err
}

func (r *StockRepository) GetByID(id uint) (*entity.Stock, error) {
	var stock entity.Stock
	err := r.db.First(&stock, id).Error
	return &stock, err
}

func (r *StockRepository) Update(stock *entity.Stock) error {
	return r.db.Model(&entity.Stock{}).Where("id = ?", stock.ID).Update("quantity", stock.Quantity).Error
}
