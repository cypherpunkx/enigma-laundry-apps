package manager

import "enigmacamp.com/enigma-laundry-apps/repository"

type RepoManager interface {
	BillRepo() repository.BillRepository
	CustomerRepo() repository.CustomerRepository
	EmployeeRepo() repository.EmployeeRepository
	ProductRepo() repository.ProductRepository
	UserRepo() repository.UserRepository
}

type repoManager struct {
	infraManager InfraManager
}

func (r *repoManager) BillRepo() repository.BillRepository {
	return repository.NewBillRepository(r.infraManager.Conn())
}

func (r *repoManager) CustomerRepo() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infraManager.Conn())
}

func (r *repoManager) EmployeeRepo() repository.EmployeeRepository {
	return repository.NewEmployeeRepository(r.infraManager.Conn())
}

func (r *repoManager) ProductRepo() repository.ProductRepository {
	return repository.NewProductRepository(r.infraManager.Conn())
}

func (r *repoManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infraManager.Conn())
}

func NewRepoManager(infra InfraManager) RepoManager {
	return &repoManager{
		infraManager: infra,
	}
}
