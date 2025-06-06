package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetAllSorted(sort_by string) ([]entity.Product, error)
	GetByID(id uint) (*entity.Product, error)
	GetByBarcode(barcode string) (*entity.Product, error)
	Create(product *entity.Product, quantity uint) error
	Update(product *entity.Product) error
	Delete(id uint) error
}
