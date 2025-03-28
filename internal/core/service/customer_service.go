package service

import (
	"errors"

	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/entity"
	"github.com/KaningNoppasin/Web-Server-Fiber/internal/core/port"
	"gorm.io/gorm"
)

type CustomerService struct {
	repo port.CustomerRepository
}

func NewCustomerService(repo port.CustomerRepository) port.CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) GetAllCustomer() ([]entity.Customer, error) {
	customers, err := s.repo.GetAll()
	if err != nil {
		return nil, ErrFailedToRetrieveCustomer
	}
	return customers, nil
}

func (s *CustomerService) GetCustomerByID(id uint) (*entity.Customer, error) {
	customer, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, ErrFailedToRetrieveCustomer
	}
	return customer, nil
}

func (s *CustomerService) GetCustomerByCardUID(card_uid string) (*entity.Customer, error) {
	customer, err := s.repo.GetByCardUID(card_uid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCustomerNotFound
		}
		return nil, ErrFailedToRetrieveCustomer
	}
	return customer, nil
}

func (s *CustomerService) CreateCustomer(customer *entity.Customer) error {
	err := s.repo.Create(customer)
	if err != nil {
		return ErrFailedToCreateCustomer
	}
	return nil
}

func (s *CustomerService) UpdateCustomer(customer *entity.Customer) error {
	_, err := s.repo.GetByID(customer.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCustomerNotFound
		}
		return ErrFailedToRetrieveCustomer
	}

	err = s.repo.Update(customer)
	if err != nil {
		return ErrFailedToCreateCustomer
	}
	return nil
}

func (s *CustomerService) DeleteCustomer(id uint) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrCustomerNotFound
		}
		return ErrFailedToRetrieveCustomer
	}

	err = s.repo.Delete(id)
	if err != nil {
		return ErrFailedToDeleteCustomer
	}
	return nil
}
