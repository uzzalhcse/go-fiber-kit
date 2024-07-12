// Http/Requests/forget_password_request.go

package authrequests

import (
	"github.com/uzzalhcse/amadeus-go/app/http/requests"
)

type ForgetPasswordRequest struct {
	// Add fields for initiating forget password process
	*requests.Validate
}
