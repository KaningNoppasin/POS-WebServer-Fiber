package db

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"gorm.io/gorm"
)

type BillRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) port.BillRepository {
	return &BillRepository{db: db}
}

func (r *BillRepository) GetAll() ([]entity.Bill, error) {
	var bills []entity.Bill
	// TODO: Change Query to r.db.Preload("Bill_Details.Product").Find(&bills).Error
	err := r.db.Preload("Bill_Details").Find(&bills).Error
	return bills, err
}

func (r *BillRepository) GetByID(id uint) (*entity.Bill, error) {
	var bill entity.Bill
	err := r.db.Preload("Bill_Details").First(&bill, id).Error
	return &bill, err
}

func (r *BillRepository) Create(bill *entity.Bill) error {
	return r.db.Create(bill).Error
}

func (r *BillRepository) Update(bill *entity.Bill) error {
	return r.db.Save(bill).Error
}

func (r *BillRepository) Delete(id uint) error {
	var bill entity.Bill
	err := r.db.Preload("Bill_Details").First(&bill, id).Error
	if err != nil {
		return err
	}

	return r.db.Select("Bill_Details").Delete(&bill).Error
}
