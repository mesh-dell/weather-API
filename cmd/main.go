package main

import (
	"github.com/mesh-dell/weather-API/internal/api"
	"github.com/mesh-dell/weather-API/internal/cache"
	"github.com/mesh-dell/weather-API/internal/config"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	cache := cache.NewCache(config)
	api.InitServer(config, &cache)
}
