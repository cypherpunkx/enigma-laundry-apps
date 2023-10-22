package model

import "time"

type Bill struct {
	ID          string        `json:"id"`
	BillDate    time.Time     `json:"billDate"`
	EntryDate   time.Time     `json:"entryDate"`
	EmployeeID  string        `json:"employeeID"`
	CustomerID  string        `json:"customerID"`
	BillDetails []BillDetails `json:"billDetails"`
}

type BillDetails struct {
	ID           string    `json:"id"`
	BillID       string    `json:"billID"`
	ProductID    string    `json:"productID"`
	ProductPrice int       `json:"productPrice"`
	Qty          int       `json:"qty"`
	FinishDate   time.Time `json:"finishDate"`
}
