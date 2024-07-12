package authrequests

import "github.com/uzzalhcse/amadeus-go/app/http/requests"

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=6"`
	*requests.Validate
}
