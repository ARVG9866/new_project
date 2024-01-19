package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/ARVG9866/new_project/internal/api/handlers"
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

	router.Get("/helloWorld", handler.Hi)
	_ = ctx

	http.ListenAndServe(":8000", router)

}
