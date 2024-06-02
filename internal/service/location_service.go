package service

import (
	"Weather-API-GoExpertPostGrad/internal/repository"
)

func GetLocationByCep(cep string) (string, error) {
	return repository.GetLocationFromViaCEP(cep)
}
