package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetWeather_HappyPath(t *testing.T) {
	_, err := GetWeather("London")
	assert.Nil(t, err, "GetWeather failed for valid location")
}

func TestGetWeather_EmptyLocation(t *testing.T) {
	_, err := GetWeather("")
	assert.NotNil(t, err, "GetWeather did not fail for empty location")
}

func TestNewTemperatureResponse_HappyPath(t *testing.T) {
	resp := NewTemperatureResponse(0)
	assert.Equal(t, 0.0, resp.TempC, "NewTemperatureResponse returned incorrect TempC")
	assert.Equal(t, 32.0, resp.TempF, "NewTemperatureResponse returned incorrect TempF")
	assert.Equal(t, 273.15, resp.TempK, "NewTemperatureResponse returned incorrect TempK")
}

func TestCelsiusToFahrenheit_HappyPath(t *testing.T) {
	assert.Equal(t, 32.0, celsiusToFahrenheit(0), "celsiusToFahrenheit returned incorrect value")
}

func TestCelsiusToKelvin_HappyPath(t *testing.T) {
	assert.Equal(t, 273.15, celsiusToKelvin(0), "celsiusToKelvin returned incorrect value")
}
