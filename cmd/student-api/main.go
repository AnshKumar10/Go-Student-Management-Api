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

	"github.com/AnshKumar10/Go-Student-Management-Api/internal/config"
	"github.com/AnshKumar10/Go-Student-Management-Api/internal/http/handlers/student"
	"github.com/AnshKumar10/Go-Student-Management-Api/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage init", slog.String("env", cfg.Env))

	router := http.NewServeMux()

	router.HandleFunc("GET /api/students", student.GetAll(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("PATCH /api/students/{id}", student.UpdateById(storage))
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteById(storage))

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("Server started")

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server Shutdown successfully")

}
