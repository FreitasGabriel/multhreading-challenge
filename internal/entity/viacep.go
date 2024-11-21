package entity

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/FreitasGabriel/multhreading-challenge/internal/dto"
)

type ViaCEPPayload struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func GetCEPFromViaCEP(cep string, cepChan chan *dto.Cep) (*dto.Cep, error) {
	var viaCEPPayload ViaCEPPayload
	request, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		fmt.Println("error to request cep", err)
		return nil, err
	}

	err = json.NewDecoder(request.Body).Decode(&viaCEPPayload)
	if err != nil {
		fmt.Println("error to decode cep", err)
		return nil, err
	}

	newCEP := dto.NewCEP(
		viaCEPPayload.Cep,
		viaCEPPayload.Uf,
		viaCEPPayload.Localidade,
		viaCEPPayload.Bairro,
		viaCEPPayload.Logradouro,
		"ViaCEP",
	)

	cepChan <- newCEP

	return newCEP, nil
}
