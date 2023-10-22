package model

type UserCredential struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsActive bool   `json:"isActive"`
}
