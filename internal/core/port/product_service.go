package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type ProductService interface {
	GetAllProduct() ([]entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	GetProductByBarcode(barcode string) (*entity.Product, error)
	CreateProduct(product *entity.Product, quantity uint) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id uint) error
}
