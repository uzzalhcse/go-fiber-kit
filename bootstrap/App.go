// bootstrap/app.go
package bootstrap

import (
	"sync"

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

	return AppInstance
}

// ConnectDB lazily connects to the database if not already connected
func (app *Application) ConnectDB() {
	go func() {
		if app.DB == nil {
			app.DB = NewDatabase(app.Config.Database)
		}
	}()
}
func (app *Application) CloseDBConnection() {
	CloseDBConnection(app.DB)
}

func (app *Application) Run(port string) error {
	return app.Listen(port)
}
