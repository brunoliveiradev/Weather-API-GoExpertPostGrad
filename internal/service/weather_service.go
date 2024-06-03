package service

import (
	"Weather-API-GoExpertPostGrad/internal/model"
	"Weather-API-GoExpertPostGrad/internal/repository"
)

func GetWeather(location string) (float64, error) {
	return repository.FetchWeather(location)
}

func NewTemperatureResponse(tempC float64) model.TemperatureResponse {
	return model.TemperatureResponse{
		TempC: tempC,
		TempF: celsiusToFahrenheit(tempC),
		TempK: celsiusToKelvin(tempC),
	}
}

func celsiusToFahrenheit(c float64) float64 {
	return c*1.8 + 32
}

func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}
