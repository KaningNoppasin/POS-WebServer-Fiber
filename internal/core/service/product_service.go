package service

import (
	"errors"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/KaningNoppasin/Web-Server-Fiber/pkg/util"
	"gorm.io/gorm"
)

type ProductService struct {
	repo port.ProductRepository
}

func NewProductService(repo port.ProductRepository) port.ProductSevice {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProduct() ([]entity.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, ErrFailedToRetrieveProduct
	}
	return products, nil
}

func (s *ProductService) GetProductByID(id uint) (*entity.Product, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, ErrFailedToRetrieveProduct
	}
	return product, nil
}

func (s *ProductService) GetProductByBarcode(barcode string) (*entity.Product, error) {
	product, err := s.repo.GetByBarcode(barcode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProductNotFound
		}
		return nil, ErrFailedToRetrieveProduct
	}
	return product, err
}

func (s *ProductService) CreateProduct(product *entity.Product) error {
	err := s.repo.Create(product)
	if err != nil {
		// if process fial need to delete image
		util.DeleteImage(product.ImagePath)

		return ErrFailedToCreateProduct
	}
	return nil
}

func (s *ProductService) UpdateProduct(product *entity.Product) error {
	existingProduct, err := s.repo.GetByID(product.ID)
	if err != nil {
		// if process fial need to delete image
		util.DeleteImage(product.ImagePath)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}
		return ErrFailedToRetrieveProduct
	}

	err = s.repo.Update(product)
	if err != nil {
		// if process fial need to delete image
		util.DeleteImage(product.ImagePath)

		return ErrFailedToCreateProduct
	}
	// if success need to delete old image
	util.DeleteImage(existingProduct.ImagePath)
	return nil
}

func (s *ProductService) DeleteProduct(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}
		return ErrFailedToRetrieveProduct
	}
	err = s.repo.Delete(id)
	if err != nil {
		return ErrFailedToDeleteProduct
	}
	return nil
}
