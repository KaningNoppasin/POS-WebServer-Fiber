package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type BillRepository interface {
	GetAll() ([]entity.Bill, error)
	GetByID(id uint) (*entity.Bill, error)
	Create(bill *entity.Bill) error
	Update(bill *entity.Bill) error
	Delete(id uint) error
}
