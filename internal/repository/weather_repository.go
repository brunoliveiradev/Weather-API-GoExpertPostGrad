package repository

import (
	"Weather-API-GoExpertPostGrad/internal/model"
	httpclient "Weather-API-GoExpertPostGrad/pkg/http"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const apiKey = "e1fece5bce574041a9f130048241703"

func FetchWeather(location string) (float64, error) {
	client := httpclient.NewClient()
	resp, err := client.Get(fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, url.QueryEscape(location)))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("failed to get weather from external API, status code: ", resp.StatusCode)
		return 0, fmt.Errorf("failed to get weather")
	}

	var weatherResponse model.WeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return 0, err
	}

	return weatherResponse.Current.TempC, nil
}
