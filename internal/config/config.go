package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	RedisUrl string
	APIKey   string
	APIBase  string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	weatherApiKey := os.Getenv("WEATHER_API_KEY")
	port := os.Getenv("PORT")
	weatherApiBase := os.Getenv("WEATHER_API_BASE")
	redisUrl := os.Getenv("REDIS_URL")

	if weatherApiKey == "" {
		return nil, fmt.Errorf("please provide a visual crossing api key")
	}

	return &Config{
		Port:     port,
		RedisUrl: redisUrl,
		APIKey:   weatherApiKey,
		APIBase:  weatherApiBase,
	}, nil
}
