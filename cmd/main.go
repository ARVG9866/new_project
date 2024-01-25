package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ARVG9866/new_project/internal/api/handlers"
	"github.com/ARVG9866/new_project/internal/app"
	"github.com/go-chi/chi/v5"
)

func main() {
	logHandlerOpt := slog.HandlerOptions{Level: slog.LevelDebug}
	logHandler := slog.NewTextHandler(os.Stdout, &logHandlerOpt)
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	ctx := context.Background()

	handler := handlers.NewHandler()
	router := chi.NewRouter()

	router.Get("/hi", handler.Hi)

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.Error("", err)
	}

	err = a.Run(router)
	if err != nil {
		logger.Error("", err)
	}
}
