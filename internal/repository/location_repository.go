package repository

import (
	"Weather-API-GoExpertPostGrad/internal/model"
	httpclient "Weather-API-GoExpertPostGrad/pkg/http"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetLocationFromViaCEP(cep string) (string, error) {
	client := httpclient.NewClient()
	resp, err := client.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("failed to get location from ViaCEP, status code: ", resp.StatusCode)
		return "", fmt.Errorf("failed to get location")
	}

	var address model.AddressResponse
	err = json.NewDecoder(resp.Body).Decode(&address)
	if address.Erro {
		return "", fmt.Errorf("zipcode not found")
	}

	if err != nil {
		return "", err
	}

	if address.Localidade == "" {
		return "", fmt.Errorf("can not find zipcode")
	}

	return address.Localidade, nil
}
