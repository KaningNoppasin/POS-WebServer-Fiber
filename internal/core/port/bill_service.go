package port

import "github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"

type BillService interface {
	GetAllBill() ([]entity.Bill, error)
	GetBillByID(id uint) (*entity.Bill, error)
	CreateBill(bill *entity.Bill) error
	UpdateBill(bill *entity.Bill) error
	DeleteBill(id uint) error
}
