// container/test_service_provider.go

package container

import (
	"github.com/uzzalhcse/amadeus-go/app/Repositories"
	"github.com/uzzalhcse/amadeus-go/app/Services"
)

// TestServiceProvider is a service provider for the TestService
type TestServiceProvider struct{}

// Register registers the TestService in the container
func (p *TestServiceProvider) Register(container *Container) {
	container.TestService = Services.NewTestService(Repositories.NewTestRepository(container.App.DB))
}
