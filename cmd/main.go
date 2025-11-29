package main

import (
	"github.com/mesh-dell/weather-API/internal/api"
	"github.com/mesh-dell/weather-API/internal/config"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	api.InitServer(config)
}
