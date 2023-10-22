package service

import (
	"testing"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var dmCustomers []model.Customer = []model.Customer{
	{
		ID:          "1",
		Name:        "Rafly",
		PhoneNumber: "087802772712",
		Address:     "Jl. Kopo Cetarip TImur 2 No. 4",
	},
}

type repoMock struct {
	mock.Mock
}

type CustomerServiceTestSuite struct {
	suite.Suite
	repoMock *repoMock
	service  CustomerService
}

func (suite *CustomerServiceTestSuite) SetupTest() {
	suite.repoMock = new(repoMock)
	suite.service = NewCustomerService(suite.repoMock)
}

func (r *repoMock) Create(payload *model.Customer) (*model.Customer, error) {
	args := r.Called(payload)

	if args[1] != nil {
		return nil, args.Error(0)
	}

	return args[1].(*model.Customer), nil
}

func (r *repoMock) List() ([]*model.Customer, error) {
	return nil, nil
}

func (r *repoMock) Update(payload *model.Customer) (*model.Customer, error) {
	return nil, nil
}

func (r *repoMock) Get(id string) (*model.Customer, error) {
	return nil, nil
}

func (r *repoMock) Delete(id string) error {
	return nil
}

func (r *repoMock) Paging(requestPaging dto.PaginationParam, params ...string) ([]*model.Customer, dto.Paging, error) {
	args := r.Called(requestPaging)

	if args[2] != nil {
		return nil, dto.Paging{}, args.Error(2)
	}

	return args[0].([]*model.Customer), args[1].(dto.Paging), nil
}

func (suite *CustomerServiceTestSuite) TestRegisterNewCustomerSuccess() {
	dmCustomer := dmCustomers[0]

	suite.repoMock.On("Create", dmCustomer).Return(nil)

	actualData, actualError := suite.service.RegisterNewCustomer(&dmCustomer)

	assert.Nil(suite.Suite.T(), actualError)
	assert.Equal(suite.Suite.T(), dmCustomer, actualData)
}

func (suite *CustomerServiceTestSuite) TestFindAllCustomerSuccess() {
	dummy := dmCustomers

	exptectedPaging := dto.Paging{
		Page:        1,
		RowsPerPage: 5,
		TotalRows:   5,
		TotalPages:  1,
	}

	requestPaging := dto.PaginationParam{
		Page:  1,
		Limit: 5,
	}

	suite.repoMock.On("Paging", requestPaging).Return(dummy, exptectedPaging, nil)

	actualProducts, actualPaging, actualError := suite.service.FindAllCustomer(requestPaging)

	assert.Equal(suite.T(), dummy, actualProducts)
	assert.Equal(suite.T(), exptectedPaging, actualPaging)
	assert.Nil(suite.T(), actualError)
}
func TestCustomerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerServiceTestSuite))
}
