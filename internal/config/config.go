package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WEATHER_API          string
	WEATHER_API_BASE_URL string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	config := &Config{
		WEATHER_API:          getEnv("WEATHER_API", true),
		WEATHER_API_BASE_URL: getEnv("WEATHER_API_BASE_URL", true),
	}

	return config
}

func getEnv(key string, required bool) string {
	env := os.Getenv(key)
	if env == "" && required {
		log.Fatalf("Variable %s required, but empty!", key)
	}

	return env
}
