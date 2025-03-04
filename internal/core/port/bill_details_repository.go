package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type Bill_DetailsRepository interface {
	GetAll() ([]entity.Bill_Details, error)
	GetByID(id uint) (*entity.Bill_Details, error)
	Create(bill_details *entity.Bill_Details) error
	Update(bill_details *entity.Bill_Details) error
	Delete(id uint) error
}
