package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Port     string
	RedisUrl string
	APIKey   string
	APIBase  string
}

func GetConfig() config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	weatherApiKey := os.Getenv("WEATHER_API_KEY")
	port := os.Getenv("PORT")
	weatherApiBase := os.Getenv("WEATHER_API_BASE")
	redisUrl := os.Getenv("REDIS_URL")

	if weatherApiKey == "" {
		log.Fatal("Please provide a visual crossing api key")
	}

	return config{
		Port:     port,
		RedisUrl: redisUrl,
		APIKey:   weatherApiKey,
		APIBase:  weatherApiBase,
	}
}
