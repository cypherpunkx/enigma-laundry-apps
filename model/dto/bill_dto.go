package dto

import (
	"time"

	"enigmacamp.com/enigma-laundry-apps/model"
)

type BillResponse struct {
	ID          string                `json:"id"`
	BillDate    time.Time             `json:"billDate"`
	EntryDate   time.Time             `json:"entryDate"`
	Employee    model.Employee        `json:"employee"`
	Customer    model.Customer        `json:"customer"`
	BillDetails []BillDetailsResponse `json:"billDetails"`
	TotalBill   int                   `json:"totalBill"`
}

type BillDetailsResponse struct {
	ID           string        `json:"id"`
	BillID       string        `json:"billID"`
	Product      model.Product `json:"product"`
	ProductPrice int           `json:"productPrice"`
	Qty          int           `json:"qty"`
	FinishDate   time.Time     `json:"finishDate"`
}
