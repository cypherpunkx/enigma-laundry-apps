package model

type Customer struct {
	ID          string `json:"id" binding:"omitempty,required"`
	Name        string `json:"name" binding:"required,alpha"`
	PhoneNumber string `json:"phoneNumber" binding:"required,numeric"`
	Address     string `json:"address" binding:"required"`
}
