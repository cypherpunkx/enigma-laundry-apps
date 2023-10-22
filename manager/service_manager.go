package manager

import "enigmacamp.com/enigma-laundry-apps/service"

type ServiceManager interface {
	EmployeeService() service.EmployeeService
	CustomerService() service.CustomerService
	ProductService() service.ProductService
	BillService() service.BillService
	UserService() service.UserService
	AuthService() service.AuthService
}

type serviceManager struct {
	repoManager RepoManager
}

func NewServiceManager(repo RepoManager) ServiceManager {
	return &serviceManager{
		repoManager: repo,
	}
}

func (s *serviceManager) EmployeeService() service.EmployeeService {
	return service.NewEmployeeService(s.repoManager.EmployeeRepo())
}

func (s *serviceManager) CustomerService() service.CustomerService {
	return service.NewCustomerService(s.repoManager.CustomerRepo())
}

func (s *serviceManager) ProductService() service.ProductService {
	return service.NewProductService(s.repoManager.ProductRepo())
}

func (s *serviceManager) BillService() service.BillService {
	return service.NewBillService(s.repoManager.BillRepo(), s.EmployeeService(), s.CustomerService(), s.ProductService())
}

func (s *serviceManager) UserService() service.UserService {
	return service.NewUserService(s.repoManager.UserRepo())
}

func (s *serviceManager) AuthService() service.AuthService {
	return service.NewAuthService(s.UserService())
}
