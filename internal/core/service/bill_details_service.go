package service

import (
	"errors"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"gorm.io/gorm"
)

type Bill_DetailsService struct {
	repo port.Bill_DetailsRepository
}

func NewBill_DetailsService(repo port.Bill_DetailsRepository) port.Bill_DetailsService {
	return &Bill_DetailsService{repo: repo}
}

func (s *Bill_DetailsService) GetAllBill_Details() ([]entity.Bill_Details, error) {
	bill_details, err := s.repo.GetAll()
	if err != nil {
		return nil, ErrFailedToRetrieveBill_Details
	}
	return bill_details, nil
}

func (s *Bill_DetailsService) GetBill_DetailsByID(id uint) (*entity.Bill_Details, error) {
	bill_detail, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBill_DetailsNotFound
		}
		return nil, ErrFailedToRetrieveBill_Details
	}
	return bill_detail, nil
}

func (s *Bill_DetailsService) CreateBill_Details(bill_details *entity.Bill_Details) error {
	err := s.repo.Create(bill_details)
	if err != nil {
		return ErrFailedToCreateBill_Details
	}
	return nil
}

func (s *Bill_DetailsService) UpdateBill_Details(bill_details *entity.Bill_Details) error {
	_, err := s.repo.GetByID(bill_details.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrBill_DetailsNotFound
		}
		return ErrFailedToRetrieveBill_Details
	}

	err = s.repo.Update(bill_details)
	if err != nil {
		return ErrFailedToCreateBill_Details
	}
	return nil
}

func (s *Bill_DetailsService) DeleteBill_Details(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrBill_DetailsNotFound
		}
		return ErrFailedToRetrieveBill_Details
	}
	err = s.repo.Delete(id)
	if err != nil {
		return ErrFailedToDeleteBill_Details
	}
	return nil
}
