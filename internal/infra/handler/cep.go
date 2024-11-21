package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/FreitasGabriel/multhreading-challenge/internal/dto"
	"github.com/FreitasGabriel/multhreading-challenge/internal/entity"
	"github.com/go-chi/chi"
)

func GetCEP(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	brasilAPIChan := make(chan *dto.Cep)
	viaCEPChan := make(chan *dto.Cep)

	go entity.GetCEPFromBrasilAPI(cep, brasilAPIChan)
	go entity.GetCEPFromViaCEP(cep, viaCEPChan)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	select {
	case cepPayload := <-viaCEPChan:
		logger.Info(fmt.Sprintf("Received message from %s: %s", cepPayload.Resource, cepPayload))
	case cepPayload := <-brasilAPIChan:
		logger.Info(fmt.Sprintf("Received message from %s: %s", cepPayload.Resource, cepPayload))
	case <-time.After(time.Second * 1):
		logger.Info("Timeout")
	}

	w.WriteHeader(http.StatusOK)
}
