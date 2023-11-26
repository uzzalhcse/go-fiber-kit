package Providers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/routes"
)

type RouteServiceProvider struct {
	*ServiceProvider
}

func (provider *RouteServiceProvider) Resister(router fiber.Router) {
	web := router.Group("")
	routes.SetUpWebRoutes(web)

	api := router.Group("/api")
	routes.SetUpApiRoutes(api)

	auth := router.Group("/api/auth")
	routes.SetUpAuthRoutes(auth)
}
