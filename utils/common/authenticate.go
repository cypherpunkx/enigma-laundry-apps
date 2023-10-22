package common

import (
	"fmt"

	"enigmacamp.com/enigma-laundry-apps/model"
)

// Fungsi untuk mengautentikasi pengguna berdasarkan username dan password
func AuthenticateUser(username, password string) (*model.UserCredential, error) {
	var users = []model.UserCredential{}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("Authentication failed")
}
