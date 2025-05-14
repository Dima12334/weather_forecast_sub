package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"weather_forecast_sub/internal/config"
	"weather_forecast_sub/internal/db"
	"weather_forecast_sub/internal/handlers"
	"weather_forecast_sub/internal/repository"
	"weather_forecast_sub/internal/server"
	"weather_forecast_sub/internal/service"
	"weather_forecast_sub/pkg/logger"
)

// Run initializes the whole application.
func Run(configDir string) {
	cfg, err := config.Init(configDir)
	if err != nil {
		log.Fatalf("failed to init configs: %v", err.Error())
		return
	}

	if err := logger.Init(cfg.Logger); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	dbConn, err := db.ConnectDB(cfg.DB)
	if err != nil {
		logger.Errorf("failed to connect to database: %v", err.Error())
		return
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			logger.Errorf("error occurred on db connection close: %s", err.Error())
		} else {
			logger.Info("db connection closed successfully")
		}
	}()

	repositories := repository.NewRepositories(dbConn)
	services := service.NewServices(
		service.Deps{
			Repos: repositories,
		},
	)
	handler := handlers.NewHandler(services)

	srv := server.NewServer(cfg, handler.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	logger.Info("server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err.Error())
	} else {
		logger.Info("server stopped successfully")
	}
}
