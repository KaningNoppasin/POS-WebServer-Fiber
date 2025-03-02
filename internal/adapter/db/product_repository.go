package db

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) port.ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetByID(id uint) (*entity.Product, error) {
	var product entity.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) GetByBarcode(barcode string) (*entity.Product, error) {
	var product entity.Product
	err := r.db.Where("product_barcode = ?", barcode).First(&product).Error
	return &product, err
}

func (r *ProductRepository) Create(product *entity.Product) error {
	err := r.db.Create(product).Error
	return err
}

func (r *ProductRepository) Update(product *entity.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}
