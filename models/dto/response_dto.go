package dto

// struct untuk pagination untuk response data (banyak dan sedikit)
type Paging struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalRows  int `json:"totalRows"`
	TotalPages int `json:"totalPages"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// response datanya satu
type SingleResponse struct {
	Status Status `json:"status"`
	Data   any    `json:"data"`
}

type PagingResponse struct {
	Status Status `json:"status"`
	Data   []any  `json:"data"`
	Paging Paging `json:"paging"`
}
