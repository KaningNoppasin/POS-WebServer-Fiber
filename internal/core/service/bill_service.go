package service

import (
	"errors"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"gorm.io/gorm"
)

type BillService struct {
	repo port.BillRepository
}

func NewBillService(repo port.BillRepository) port.BillService {
	return &BillService{repo: repo}
}

func (s *BillService) GetAllBill() ([]entity.Bill, error) {
	bills, err := s.repo.GetAll()
	if err != nil {
		return nil, ErrFailedToRetrieveBill
	}
	return bills, nil
}

func (s *BillService) GetBillByID(id uint) (*entity.Bill, error) {
	bill, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBillNotFound
		}
		return nil, ErrFailedToRetrieveBill
	}
	return bill, nil
}

func (s *BillService) CreateBill(createBillRequest *entity.CreateBillRequest) error {
	var total_amount uint
	for _, req_bill_detail := range createBillRequest.Bill_Details {
		// get product price from database
		product_price, err := s.repo.GetProductPriceByID(req_bill_detail.ProductID)
		if err != nil {
			return err
		}
		total_price := product_price * req_bill_detail.Quantity
		total_amount += total_price

		// create bill detail
		bill_detail := entity.Bill_Details{
			BillID:    req_bill_detail.BillID,
			ProductID: req_bill_detail.ProductID,
			Quantity:  req_bill_detail.Quantity,
			Total:     total_price,
		}
		err = s.repo.CreateBillDetail(&bill_detail)
		if err != nil {
			return err
		}
	}
	// create bill
	bill := entity.Bill{
		CustomerID:  createBillRequest.CustomerID,
		TotalAmount: total_amount,
	}
	return s.repo.Create(&bill)
}

func (s *BillService) UpdateBill(bill *entity.Bill) error {
	_, err := s.repo.GetByID(bill.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrBillNotFound
		}
		return ErrFailedToRetrieveBill
	}

	err = s.repo.Update(bill)
	if err != nil {
		return ErrFailedToCreateBill
	}
	return nil
}

func (s *BillService) DeleteBill(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrBillNotFound
		}
		return ErrFailedToRetrieveBill
	}
	err = s.repo.Delete(id)
	if err != nil {
		return ErrFailedToDeleteBill
	}
	return nil
}
