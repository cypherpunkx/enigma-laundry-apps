package service

import (
	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/repository"
)

type ProductService interface {
	FindAllProduct(requesPaging dto.PaginationParam, params ...string) ([]*model.Product, dto.Paging, error)
	RegisterNewProduct(payload *model.Product) (*model.Product, error)
	GetProductByID(id string) (*model.Product, error)
	UpdateProductByID(payload *model.Product) (*model.Product, error)
	DeleteProductByID(id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) FindAllProduct(requesPaging dto.PaginationParam, params ...string) ([]*model.Product, dto.Paging, error) {
	return s.repo.Paging(requesPaging, params...)
}

func (s *productService) RegisterNewProduct(payload *model.Product) (*model.Product, error) {
	return s.repo.Create(payload)
}

func (s *productService) UpdateProductByID(payload *model.Product) (*model.Product, error) {
	return s.repo.Update(payload)
}

func (s *productService) GetProductByID(id string) (*model.Product, error) {
	return s.repo.Get(id)
}

func (s *productService) DeleteProductByID(id string) error {
	return s.repo.Delete(id)
}
