// container/container.go

package container

import (
	"github.com/uzzalhcse/amadeus-go/app/Services"
	"github.com/uzzalhcse/amadeus-go/bootstrap"
	"sync"
)

// Container is a simple service container
type Container struct {
	App         *bootstrap.Application
	TestService *Services.TestService
	// Add other services as needed
	// ...
}

// GetContainer returns the singleton instance of the service container
func GetContainer(app *bootstrap.Application) *Container {
	once.Do(func() {
		containerInstance = initContainer(app)
	})
	return containerInstance
}

// initContainer initializes the container with all services
func initContainer(app *bootstrap.Application) *Container {
	container := &Container{
		App: app,
		// Initialize other services
		// ...
	}

	// Register services using service providers
	for _, provider := range serviceProviders {
		provider.Register(container)
	}

	return container
}

// RegisterServiceProvider registers a new service provider
func RegisterServiceProvider(provider ServiceProvider) {
	serviceProviders = append(serviceProviders, provider)
}

var containerInstance *Container
var once sync.Once
var serviceProviders []ServiceProvider
