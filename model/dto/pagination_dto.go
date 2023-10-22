package dto

// Ini buat paging di taruh di parameter
type PaginationParam struct {
	Page   int
	Offset int
	Limit  int
}

// ini buat paging di taruh di return
type PaginationQuery struct {
	Page int
	Take int
	Skip int
}

// ini buat di taruh di response
type Paging struct {
	Page        int `json:"omitempty"`
	RowsPerPage int `json:"omitempty"`
	TotalRows   int `json:"omitempty"`
	TotalPages  int `json:"omitempty"`
}

// example pagination -> product 100
// Paging {Page: 1, RowsPerPage: 10, TotalRows: 100, TotalPages: 10}
