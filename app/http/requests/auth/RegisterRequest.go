// Http/Requests/register_request.go

package authrequests

import "github.com/uzzalhcse/go-fiber-kit/app/http/requests"

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	*requests.Validate
}
