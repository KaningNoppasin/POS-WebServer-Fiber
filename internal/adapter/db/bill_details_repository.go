package db

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"gorm.io/gorm"
)

type Bill_DetailsRepository struct {
	db *gorm.DB
}

func NewBill_DetailsRepository(db *gorm.DB) port.Bill_DetailsRepository {
	return &Bill_DetailsRepository{db: db}
}

func (r *Bill_DetailsRepository) GetAll() ([]entity.Bill_Details, error) {
	var bill_details []entity.Bill_Details
	err := r.db.Preload("Product.Stock").Find(&bill_details).Error
	return bill_details, err
}

func (r *Bill_DetailsRepository) GetByID(id uint) (*entity.Bill_Details, error) {
	var bill_detail entity.Bill_Details
	err := r.db.Preload("Product.Stock").First(&bill_detail, id).Error
	return &bill_detail, err
}

func (r *Bill_DetailsRepository) Create(bill_details *entity.Bill_Details) error {
	return r.db.Create(bill_details).Error
}

func (r *Bill_DetailsRepository) Update(bill_details *entity.Bill_Details) error {
	return r.db.Save(bill_details).Error
}

func (r *Bill_DetailsRepository) Delete(id uint) error {
	var bill_details entity.Bill_Details
	err := r.db.Preload("Product.Stock").First(&bill_details, id).Error
	if err != nil {
		return err
	}

	return r.db.Delete(&bill_details).Error
}
