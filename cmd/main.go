package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-boilerplate/internal/api/controller"
	"go-boilerplate/internal/api/router"
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/repository"
	"go-boilerplate/internal/service"
	"go-boilerplate/pkg/logger"
)

func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Initialize Logger
	appLogger := logger.NewLogger()
	appLogger.Info("Starting the application...")

	// 3. Initialize Hub (Repositories, Services, Controllers)
	userRepo := repository.NewUserRepository()
	userSvc := service.NewUserService(userRepo)

	healthCtrl := controller.NewHealthController()
	userCtrl := controller.NewUserController(userSvc)

	// 4. Initialize Router
	r := router.SetupRouter(healthCtrl, userCtrl)

	// 5. Start Server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful Shutdown logic
	go func() {
		appLogger.Info(fmt.Sprintf("Server is running on port %s", cfg.Port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", cfg.Port, err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	appLogger.Info("Shutting down server...")
	// You can add database disconnection logic here
}
