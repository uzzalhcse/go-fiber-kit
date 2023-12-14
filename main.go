package main

import (
	"fmt"
	"github.com/uzzalhcse/amadeus-go/app/exceptions"
	"github.com/uzzalhcse/amadeus-go/bootstrap"
	"github.com/uzzalhcse/amadeus-go/routes"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	app.ConnectDB()

	// Register routes
	routes.RegisterRoutes(app.App)

	// Launch the application in a goroutine
	go startApplication(app)

	// Graceful shutdown
	app.GracefulShutdown(func() {
		shutdownApplication(app)
	})
}
func startApplication(app *bootstrap.Application) {
	port := ":" + app.Config.App.Port
	if err := app.Run(port); err != nil {
		exceptions.PanicIfNeeded(err.Error())
	}
}

func shutdownApplication(app *bootstrap.Application) {
	if err := app.Shutdown(); err != nil {
		fmt.Println("Error during shutdown:", err.Error())
	}

	app.CloseDBConnection()
}
