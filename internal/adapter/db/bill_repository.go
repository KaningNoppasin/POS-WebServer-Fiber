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
	err := r.db.Preload("Bill_Details.Product.Stock").Find(&bills).Error
	return bills, err
}

func (r *BillRepository) GetByID(id uint) (*entity.Bill, error) {
	var bill entity.Bill
	err := r.db.Preload("Bill_Details.Product.Stock").First(&bill, id).Error
	return &bill, err
}

func (r *BillRepository) GetProductPriceByID(product_id uint) (uint, error) {
	var product_price uint
	err := r.db.Model(&entity.Product{}).Select("price").Where("id = ?", product_id).Scan(&product_price).Error
	return product_price, err
}

func (r *BillRepository) Create(bill *entity.Bill) error {
	return r.db.Create(bill).Error
}

func (r *BillRepository) CreateBillDetail(bill_details *entity.Bill_Details) error {
	return r.db.Create(bill_details).Error
}

func (r *BillRepository) Update(bill *entity.Bill) error {
	return r.db.Save(bill).Error
}

func (r *BillRepository) Delete(id uint) error {
	var bill entity.Bill
	err := r.db.Preload("Bill_Details.Product.Stock").First(&bill, id).Error
	if err != nil {
		return err
	}

	return r.db.Select("Bill_Details").Delete(&bill).Error
}
