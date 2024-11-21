package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/FreitasGabriel/multhreading-challenge/internal/infra/handler"
	"github.com/go-chi/chi"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting application")

	r := chi.NewRouter()
	r.Get("/cep/{cep}", handler.GetCEP)
	http.ListenAndServe(":3000", r)
}
