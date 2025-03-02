package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type ProductSevice interface {
	GetAllProduct() ([]entity.Product, error)
	GetProductByID(id uint) (*entity.Product, error)
	GetProductByBarcode(barcode string) (*entity.Product, error)
	CreateProduct(product *entity.Product) error
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id uint) error
}
