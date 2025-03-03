package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type CustomerService interface {
	GetAllCustomer() ([]entity.Customer, error)
	GetCustomerByID(id uint) (*entity.Customer, error)
	CreateCustomer(customer *entity.Customer) error
	UpdateCustomer(customer *entity.Customer) error
	DeleteCustomer(id uint) error
}
