package exception

import "errors"

var (
	//! ERROR
	ErrNotFound          = errors.New("data not found")
	ErrCodeAlreadyExist  = errors.New("code already exist")
	ErrFailedCreate      = errors.New("failed to create data")
	ErrFailedUpdate      = errors.New("failed to update data")
	ErrTitleAlreadyExist = errors.New("title already exist")
	ErrInvalidPage       = errors.New("invalid page")
	ErrInvalidPerPage    = errors.New("invalid per page")
)
