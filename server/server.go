package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/bootstrap"
)

type Server struct {
	*fiber.App
	*bootstrap.Application
}

func NewServer(app *bootstrap.Application) *Server {
	server := &Server{
		App:         fiber.New(),
		Application: app,
	}
	return server
}

func (s *Server) Run(port string) error {
	return s.Listen(port)
}
