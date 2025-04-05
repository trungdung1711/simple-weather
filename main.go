/*
Copyright © 2025 zun HERE trungdunglebui17112004@gmail.com
*/
package main

import (
	"github.com/trungdung1711/simple-weather/cmd"
	"github.com/trungdung1711/simple-weather/internal/client"
	"github.com/trungdung1711/simple-weather/internal/config"
)

func main() {
	config := config.LoadConfig()
	cmd.Execute()
	client.FetchWeather(config)
}
