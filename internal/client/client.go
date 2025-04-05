package client

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/trungdung1711/simple-weather/internal/config"
)

func FetchWeather(config *config.Config) {
	queryParam := map[string]string{
		"key": config.WEATHER_API,
		"q":   "Tokyo",
	}

	client := resty.New().
		SetBaseURL(config.WEATHER_API_BASE_URL)
	resp, err := client.R().
		SetQueryParams(queryParam).
		Get("/current.json")

	if err != nil {
		log.Fatalf("Error when fetching weather: %e", err)
	}
	fmt.Println(resp.String())
}
