package db

import (
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) port.CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) GetAll() ([]entity.Customer, error) {
	var customers []entity.Customer
	err := r.db.Preload("Bill").Find(&customers).Error
	return customers, err
}

func (r *CustomerRepository) GetByID(id uint) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Preload("Bill").First(&customer, id).Error
	return &customer, err
}

func (r *CustomerRepository) GetByCardUID(card_uid string) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Preload("Bill").Where("card_uid = ?", card_uid).First(&customer).Error
	return &customer, err
}

func (r *CustomerRepository) GetByPhone(phone string) (*entity.Customer, error) {
	var customer entity.Customer
	err := r.db.Preload("Bill").Where("phone = ?", phone).First(&customer).Error
	return &customer, err
}

func (r *CustomerRepository) Create(customer *entity.Customer) error {
	return r.db.Create(customer).Error
}

func (r *CustomerRepository) Update(customer *entity.Customer) error {
	return r.db.Save(customer).Error
}

func (r *CustomerRepository) Delete(id uint) error {
	var customer entity.Customer
	err := r.db.Preload("Bill").First(&customer, id).Error
	if err != nil {
		return err
	}

	return r.db.Select("Bill").Delete(&customer).Error
}
