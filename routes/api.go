package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/app/Http/Controllers"
)

// func (r *RouteService) Api(router fiber.Router) {
//
//		testRepo := Repositories.NewTestRepository(r.DB)
//		testService := Services.NewTestService(testRepo)
//		testController := Controllers.NewTestController(testService)
//		router.Get("/", testController.Test)
//		router.Get("/api/test", testController.GetAllHandler)
//	}
func (r *RouteService) Api(router fiber.Router) {
	testController := Controllers.NewTestController(r.Container)

	router.Get("/", testController.Test)
	router.Get("/api/test", testController.GetAllHandler)
}
