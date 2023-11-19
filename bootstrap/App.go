package bootstrap

import (
	"github.com/uzzalhcse/amadeus-go/config"
	"gorm.io/gorm"
)

// Application App struct holds the Fiber application instance and the database connection instance
type Application struct {
	//*fiber.App
	DB     *gorm.DB
	Config *config.Config
}

// App initializes a new App instance
func App() *Application {
	app := &Application{nil, config.NewConfig()}

	// Setup database connection

	app.DB = NewDatabase(app.Config.Database)

	return app
}

func (app *Application) CloseDBConnection() {
	CloseDBConnection(app.DB)
}
