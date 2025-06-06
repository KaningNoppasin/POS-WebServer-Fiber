package service

import (
	"errors"
	"strings"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"github.com/KaningNoppasin/Web-Server-Fiber/pkg/util"
	"gorm.io/gorm"
)

type ProductService struct {
	repo port.ProductRepository
}

func NewProductService(repo port.ProductRepository) port.ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProduct() ([]entity.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, ErrFailedToRetrieveProduct
	}
	return products, nil
}

func (s *ProductService) GetAllProductSorted(sort_by string) ([]entity.Product, error) {
	if !strings.Contains(sort_by, "-") {
		return nil, ErrBadRequest
	}

	split := strings.Split(sort_by, "-")
	fieldName := split[0]
	sortOrder := split[1]
	if fieldName == "" || sortOrder == "" {
		return nil, ErrBadRequest
	}

	products, err := s.repo.GetAllSorted(sort_by)
	if err != nil {
		if errors.Is(err, ErrBadRequest) {
			return nil, ErrBadRequest
		}
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

func (s *ProductService) CreateProduct(product *entity.Product, quantity uint) error {
	err := s.repo.Create(product, quantity)
	if err != nil {
		// if process fail need to delete image
		util.DeleteImage(product.ImagePath)

		return ErrFailedToCreateProduct
	}
	return nil
}

func (s *ProductService) UpdateProduct(product *entity.Product) error {
	existingProduct, err := s.repo.GetByID(product.ID)
	if err != nil {
		// if process fail need to delete image
		util.DeleteImage(product.ImagePath)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrProductNotFound
		}
		return ErrFailedToRetrieveProduct
	}

	isNewPicture := existingProduct.ImagePath != product.ImagePath
	isChangeToDefaultImage := isNewPicture && (product.ImagePath == util.DefaultImage)

	// When didn't upload image when update product Just use old image
	if isChangeToDefaultImage {
		product.ImagePath = existingProduct.ImagePath
	}

	err = s.repo.Update(product)
	if err != nil {
		// if process fail need to delete image
		if !isChangeToDefaultImage {
			util.DeleteImage(existingProduct.ImagePath)
		}

		return ErrFailedToCreateProduct
	}
	// if success need to delete old image
	if !isChangeToDefaultImage {
		util.DeleteImage(existingProduct.ImagePath)
	}
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
