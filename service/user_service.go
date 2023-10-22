package service

import (
	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterNewUser(payload model.UserCredential) error
	FindAllUser() ([]*model.UserCredential, error)
	FindByUsername(username string) (*model.UserCredential, error)
	FindByUsernamePassword(username string, password string) (*model.UserCredential, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterNewUser(payload model.UserCredential) error {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	payload.Password = string(bytes)

	return s.repo.Create(&payload)
}

func (s *userService) FindAllUser() ([]*model.UserCredential, error) {
	return s.repo.List()
}

func (s *userService) FindByUsername(username string) (*model.UserCredential, error) {
	return s.repo.GetUsername(username)
}

func (s *userService) FindByUsernamePassword(username string, password string) (*model.UserCredential, error) {

	return s.repo.GetByUsernamePassword(username, password)
}
