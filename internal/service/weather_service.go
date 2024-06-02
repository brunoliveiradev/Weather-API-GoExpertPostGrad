package service

import (
	"Weather-API-GoExpertPostGrad/internal/model"
	"Weather-API-GoExpertPostGrad/internal/repository"
	"net/url"
)

func GetWeather(location string) (float64, error) {
	cityEncoded := url.QueryEscape(location)
	return repository.FetchWeather(cityEncoded)
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
