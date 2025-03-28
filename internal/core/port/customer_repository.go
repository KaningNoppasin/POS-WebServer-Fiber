package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type CustomerRepository interface {
	GetAll() ([]entity.Customer, error)
	GetByID(id uint) (*entity.Customer, error)
	GetByCardUID(card_uid string) (*entity.Customer, error)
	Create(customer *entity.Customer) error
	Update(customer *entity.Customer) error
	Delete(id uint) error
}
