package dto

type ErrorResponse struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}
