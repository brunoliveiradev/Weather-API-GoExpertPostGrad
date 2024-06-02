package handler

import (
	"Weather-API-GoExpertPostGrad/internal/service"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func HandleGetTemperatureByCEP(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	address, err := service.GetLocationByCep(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	tempC, err := service.GetWeather(address)
	if err != nil {
		http.Error(w, "failed to get weather", http.StatusNotFound)
		return
	}

	response := service.NewTemperatureResponse(tempC)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
