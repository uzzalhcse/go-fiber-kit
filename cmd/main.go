package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/uzzalhcse/amadeus-go/app/Providers"
	"github.com/uzzalhcse/amadeus-go/app/exceptions"
	"github.com/uzzalhcse/amadeus-go/bootstrap"
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
	app.ConnectDB()

	provider := Providers.RouteServiceProvider{}
	provider.Resister(app.App)

	go func() {
		if err := app.Run(":" + app.Config.App.Port); !errors.Is(err, http.ErrServerClosed) {
			exceptions.PanicIfNeeded(fmt.Errorf("error occurred while running http server"))
		}
	}()

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	log.Println("Server is shutting down...")

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Printf("Failed to stop server: %v", err)
		exceptions.PanicIfNeeded(fmt.Errorf("failed to stop server: %v", err))
	} else {
		log.Println("Server stopped gracefully")
	}
}
