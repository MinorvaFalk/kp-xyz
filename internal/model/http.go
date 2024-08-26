package model

type HTTPResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type HTTPPaginationRequest struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}
