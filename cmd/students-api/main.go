package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/geekyharsh05/students-api/internal/config"
	"github.com/geekyharsh05/students-api/internal/http/handlers/student"
	"github.com/geekyharsh05/students-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// setup database
	storage, err := sqlite.New(cfg)

	slog.Info("Storage initialized successfully", slog.String("env", cfg.Env))

	if err != nil {
		log.Fatal("Failed to setup database", slog.String("error", err.Error()))
	}

	// setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))

	// setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}
	
	slog.Info("Starting server", slog.String("address", cfg.HTTPServer.Address))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed To Start Server")
		}
	}()

	<-done

	slog.Info("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("Failed to shutdown the server gracefully", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully")
}
