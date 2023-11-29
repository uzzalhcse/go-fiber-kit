// controllers/BaseController.go

package Controllers

import "github.com/uzzalhcse/amadeus-go/bootstrap"

// BaseController contains the application instance
type BaseController struct {
	*bootstrap.Application
}

// NewBaseController initializes a new BaseController
func NewBaseController() *BaseController {
	return &BaseController{bootstrap.App()}
}
