package service

import (
	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/repository"
	"enigmacamp.com/enigma-laundry-apps/utils/exception"
)

type CustomerService interface {
	FindAllCustomer(requesPaging dto.PaginationParam, params ...string) ([]*model.Customer, dto.Paging, error)
	RegisterNewCustomer(payload *model.Customer) (*model.Customer, error)
	GetCustomerByID(id string) (*model.Customer, error)
	UpdateCustomerByID(payload *model.Customer) (*model.Customer, error)
	DeleteCustomerByID(id string) error
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) RegisterNewCustomer(payload *model.Customer) (*model.Customer, error) {
	return s.repo.Create(payload)
}

func (s *customerService) FindAllCustomer(requesPaging dto.PaginationParam, params ...string) ([]*model.Customer, dto.Paging, error) {
	return s.repo.Paging(requesPaging, params...)
}

func (s *customerService) UpdateCustomerByID(payload *model.Customer) (*model.Customer, error) {
	_, err := s.repo.Get(payload.ID)

	if err != nil {
		return nil, exception.ErrNotFound
	}

	return s.repo.Update(payload)
}

func (s *customerService) GetCustomerByID(id string) (*model.Customer, error) {
	return s.repo.Get(id)
}

func (s *customerService) DeleteCustomerByID(id string) error {
	return s.repo.Delete(id)
}
