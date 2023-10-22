package service

import (
	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/repository"
)

type EmployeeService interface {
	FindAllEmployee(requesPaging dto.PaginationParam, params ...string) ([]*model.Employee, dto.Paging, error)
	RegisterNewEmployee(payload *model.Employee) (*model.Employee, error)
	GetEmployeeByID(id string) (*model.Employee, error)
	UpdateEmployeeByID(payload *model.Employee) (*model.Employee, error)
	DeleteEmployeeByID(id string) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo: repo}
}

func (e *employeeService) FindAllEmployee(requesPaging dto.PaginationParam, params ...string) ([]*model.Employee, dto.Paging, error) {
	return e.repo.Paging(requesPaging, params...)
}

func (e *employeeService) RegisterNewEmployee(payload *model.Employee) (*model.Employee, error) {
	return e.repo.Create(payload)
}

func (e *employeeService) UpdateEmployeeByID(payload *model.Employee) (*model.Employee, error) {
	return e.repo.Update(payload)
}

func (e *employeeService) GetEmployeeByID(id string) (*model.Employee, error) {
	return e.repo.Get(id)
}

func (e *employeeService) DeleteEmployeeByID(id string) error {
	return e.repo.Delete(id)
}
