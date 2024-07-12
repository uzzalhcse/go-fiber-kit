// controllers/BaseController.go

package controllers

import "github.com/uzzalhcse/go-fiber-kit/bootstrap"

// BaseController contains the application instance
type BaseController struct {
	*bootstrap.Application
}

// NewBaseController initializes a new BaseController
func NewBaseController() *BaseController {
	return &BaseController{bootstrap.App()}
}
