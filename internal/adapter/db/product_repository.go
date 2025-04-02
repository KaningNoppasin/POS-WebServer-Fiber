package db

import (
	"strings"

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
	err := r.db.Preload("Stock").Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetAllSorted(sort_by string) ([]entity.Product, error) {
	var products []entity.Product
	// change product_name-desc to product_name desc
	sort_by = strings.Replace(sort_by, "-", " ", -1)
	err := r.db.Preload("Stock").Order(sort_by).Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetByID(id uint) (*entity.Product, error) {
	var product entity.Product
	err := r.db.Preload("Stock").First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) GetByBarcode(barcode string) (*entity.Product, error) {
	var product entity.Product
	err := r.db.Preload("Stock").Where("product_barcode = ?", barcode).First(&product).Error
	return &product, err
}

func (r *ProductRepository) Create(product *entity.Product, quantity uint) error {
	err := r.db.Create(product).Error
	if err != nil {
		return err
	}

	// Create Stock that Quantity initial 0
	err = r.db.Create(&entity.Stock{
		ProductID: product.ID,
		Quantity:  quantity,
	}).Error
	return err
}

func (r *ProductRepository) Update(product *entity.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	var product entity.Product
	err := r.db.Preload("Stock").First(&product, id).Error
	if err != nil {
		return err
	}

	return r.db.Select("Stock").Delete(&product).Error
}
