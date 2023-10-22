package repository

import (
	"database/sql"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/utils/constant"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(payload *model.UserCredential) error
	List() ([]*model.UserCredential, error)
	GetUsername(username string) (*model.UserCredential, error)
	GetByUsernamePassword(username string, password string) (*model.UserCredential, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(payload *model.UserCredential) error {
	_, err := r.db.Exec(constant.USER_CREATE, payload.ID, payload.Username, payload.Password, payload.IsActive)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) List() ([]*model.UserCredential, error) {
	var users []*model.UserCredential

	rows, err := r.db.Query(constant.USER_LIST)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.UserCredential

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.IsActive)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r *userRepository) GetUsername(username string) (*model.UserCredential, error) {
	var user model.UserCredential

	err := r.db.QueryRow(constant.USER_USERNAME_GET, username, true).Scan(&user.ID, &user.Username, &user.Password, &user.IsActive)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByUsernamePassword(username string, password string) (*model.UserCredential, error) {
	user, err := r.GetUsername(username)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return user, nil
}
