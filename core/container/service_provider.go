// container/service_provider.go

package container

// ServiceProvider is an interface for service providers
type ServiceProvider interface {
	Register(container *Container)
}
