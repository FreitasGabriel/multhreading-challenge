package entity

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FreitasGabriel/multhreading-challenge/internal/dto"
)

type BrasilAPICEPPayload struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func GetCEPFromBrasilAPI(cep string, cepChan chan *dto.Cep) (*dto.Cep, error) {
	var payload BrasilAPICEPPayload
	request, err := http.Get(fmt.Sprintf(
		"https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		fmt.Println("error to request cep", err)
		return nil, err
	}

	err = json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		fmt.Println("error to decode cep", err)
		return nil, err
	}

	newCEP := dto.NewCEP(
		payload.Cep,
		payload.State,
		payload.City,
		payload.Neighborhood,
		payload.Street,
		"BrasilAPI",
	)

	cepChan <- newCEP

	return newCEP, nil
}
