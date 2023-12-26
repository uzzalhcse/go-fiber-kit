// bootstrap/app.go
package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/uzzalhcse/amadeus-go/config"
	"gorm.io/gorm"
)

var Once sync.Once
var AppInstance *Application

// Application App struct holds the Fiber application instance and the database connection instance
type Application struct {
	*fiber.App
	DB     *gorm.DB
	Config *config.Config
}

// App initializes a new App instance (Singleton)
func App() *Application {
	Once.Do(func() {
		AppInstance = &Application{fiber.New(), nil, config.NewConfig()}
	})
	AppInstance = &Application{fiber.New(), nil, config.NewConfig()}

	return AppInstance
}

// ConnectDB lazily connects to the database if not already connected
func (app *Application) ConnectDB() {
	if app.DB == nil {
		app.DB = NewDatabase(app.Config.Database)
	}
}

// ConnectDBAsync lazily connects to the database if not already connected
func (app *Application) ConnectDBAsync() {
	go func() {
		if app.DB == nil {
			app.DB = NewDatabase(app.Config.Database)
		}
	}()
}

func (app *Application) GetDB() *gorm.DB {
	return app.DB
}
func (app *Application) CloseDBConnection() {
	CloseDBConnection(app.DB)
}

func (app *Application) Run(port string) error {
	return app.Listen(port)
}

func (app *Application) GracefulShutdown(cb func()) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		defer cancel() // Ensure cancellation when the goroutine exits

		sig := <-sigs
		fmt.Printf("Received %v signal. Initiating shutdown...\n", sig)
	}()

	// Wait for signal or context cancellation
	select {
	case <-sigs:
		// Signal received, proceed with cleanup
		fmt.Println("Shutting down gracefully...")
	case <-ctx.Done():
		// Context canceled, no need to handle the signal
		fmt.Println("Shutdown initiated by the application. Performing cleanup...")
	}

	cb()

	fmt.Println("Shutdown complete. Goodbye!")
}
