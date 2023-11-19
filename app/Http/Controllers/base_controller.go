// controllers/base_controller.go

package Controllers

import (
	"github.com/uzzalhcse/amadeus-go/core/container"
)

// BaseController defines a base controller with access to the service container
type BaseController struct {
	*container.Container
}

// NewBaseController creates a new instance of the base controller
func NewBaseController(container *container.Container) *BaseController {
	return &BaseController{container}
}

// Add any other common methods or properties you want to include in all controllers
