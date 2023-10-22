package service

import (
	"fmt"
	"strings"
	"time"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/repository"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"enigmacamp.com/enigma-laundry-apps/utils/exception"
)

type BillService interface {
	RegisterNewBill(payload *model.Bill) error
	ListTransactions(requestPaging dto.PaginationParam) ([]*dto.BillResponse, dto.Paging, error)
	FindBillByID(id string) (*dto.BillResponse, error)
}

type billService struct {
	repo            repository.BillRepository
	employeeService EmployeeService
	customerService CustomerService
	productService  ProductService
}

func NewBillService(repo repository.BillRepository, employee EmployeeService, customer CustomerService, product ProductService) BillService {
	return &billService{
		repo:            repo,
		employeeService: employee,
		customerService: customer,
		productService:  product,
	}
}

func (s *billService) RegisterNewBill(payload *model.Bill) error {
	customer, err := s.customerService.GetCustomerByID(payload.CustomerID)

	if err != nil {
		return exception.ErrNotFound
	}

	employee, err := s.employeeService.GetEmployeeByID(payload.EmployeeID)

	if err != nil {
		return exception.ErrNotFound
	}

	newBillDetail := make([]model.BillDetails, 0, len(payload.BillDetails))

	for _, billDetail := range payload.BillDetails {
		product, err := s.productService.GetProductByID(billDetail.ProductID)

		if err != nil {
			return fmt.Errorf("product with id %s is not found", billDetail.ID)
		}

		if strings.ToLower(product.Name) == "reguler" {
			billDetail.FinishDate = time.Now().Add(time.Hour * 24 * 3)
		} else {
			billDetail.FinishDate = time.Now().Add(time.Hour * 24)
		}

		billDetail.ID = common.GenerateUUID()
		billDetail.BillID = payload.ID
		billDetail.ProductID = product.ID
		billDetail.ProductPrice = product.Price
		newBillDetail = append(newBillDetail, billDetail)
	}

	payload.BillDate = time.Now()
	payload.EntryDate = time.Now()
	payload.CustomerID = customer.ID
	payload.EmployeeID = employee.ID
	payload.BillDetails = newBillDetail

	return s.repo.Create(*payload)
}

func (s *billService) ListTransactions(requestPaging dto.PaginationParam) ([]*dto.BillResponse, dto.Paging, error) {
	return s.repo.List(requestPaging)
}

func (s *billService) FindBillByID(id string) (*dto.BillResponse, error) {

	bill, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	employee, err := s.customerService.GetCustomerByID(bill.CustomerID)

	if err != nil {
		return nil, err
	}

	customer, err := s.employeeService.GetEmployeeByID(bill.EmployeeID)

	if err != nil {
		return nil, err
	}

	billDetails, err := s.repo.GetBillDetailByBill(bill.ID)

	if err != nil {
		return nil, err
	}

	var billDetailsResponse []dto.BillDetailsResponse

	var total int
	for _, billDetail := range billDetails {
		product, err := s.productService.GetProductByID(billDetail.ProductID)

		if err != nil {
			return nil, err
		}

		billDetailsResponse = append(billDetailsResponse, dto.BillDetailsResponse{
			ID:           billDetail.BillID,
			BillID:       bill.ID,
			Product:      *product,
			ProductPrice: billDetail.ProductPrice,
			Qty:          billDetail.Qty,
			FinishDate:   billDetail.FinishDate,
		})

		total += product.Price * billDetail.Qty
	}

	return &dto.BillResponse{
		ID:          bill.ID,
		BillDate:    bill.BillDate,
		EntryDate:   bill.EntryDate,
		Employee:    model.Employee(*employee),
		Customer:    model.Customer(*customer),
		BillDetails: billDetailsResponse,
		TotalBill:   total,
	}, nil
}
