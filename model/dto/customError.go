package dto

type ErrNotFound struct {
	Message string
}

func (err *ErrNotFound) Error() string {
	err.Message = "Id Not Found!"

	return err.Message
}
