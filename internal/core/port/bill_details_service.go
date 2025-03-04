package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type Bill_DetailsService interface {
	GetAllBill_Details() ([]entity.Bill_Details, error)
	GetBill_DetailsByID(id uint) (*entity.Bill_Details, error)
	CreateBill_Details(bill_details *entity.Bill_Details) error
	UpdateBill_Details(bill_details *entity.Bill_Details) error
	DeleteBill_Details(id uint) error
}
