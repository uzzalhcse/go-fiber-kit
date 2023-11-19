// main.go

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/uzzalhcse/amadeus-go/app/exceptions"
	"github.com/uzzalhcse/amadeus-go/bootstrap"
	"github.com/uzzalhcse/amadeus-go/core/container"
	"github.com/uzzalhcse/amadeus-go/routes"
	"github.com/uzzalhcse/amadeus-go/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()

	// Register service providers
	container.RegisterServiceProvider(&container.TestServiceProvider{})
	// Register more service providers as needed

	// Initialize the service container
	containerInstance := container.GetContainer(app)

	srv := server.NewServer(app)
	route := routes.NewRoutesApp(srv, containerInstance)
	route.RegisterRoute()

	go func() {
		if err := srv.Run(":" + app.Config.App.Port); !errors.Is(err, http.ErrServerClosed) {
			exceptions.PanicIfNeeded(fmt.Errorf("error occurred while running http server: %s\n", err.Error()))
		}
	}()

	log.Printf("Server is running on port %s\n", app.Config.App.Port)

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.ShutdownWithContext(ctx); err != nil {
		exceptions.PanicIfNeeded(fmt.Errorf("failed to stop server: %v", err))
	}
}
