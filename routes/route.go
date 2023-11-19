package routes

import (
	"github.com/uzzalhcse/amadeus-go/core/container"
	"github.com/uzzalhcse/amadeus-go/server"
)

type RouteService struct {
	*server.Server
	*container.Container
}

func NewRoutesApp(srv *server.Server, container *container.Container) *RouteService {
	return &RouteService{srv, container}
}

func (r *RouteService) RegisterRoute() {
	v1 := r.Group("/api/v1")
	{
		r.Api(v1)
	}
	r.Web(v1)

}
