package response

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

type ErrorDetail struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
	Title  string `json:"title"`
}
