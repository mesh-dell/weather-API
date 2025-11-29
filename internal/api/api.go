package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mesh-dell/weather-API/internal/cache"
	"github.com/mesh-dell/weather-API/internal/config"
	"github.com/mesh-dell/weather-API/internal/handlers"
	"github.com/mesh-dell/weather-API/internal/middleware"
	"github.com/mesh-dell/weather-API/internal/services"
)

func InitServer(config *config.Config, cache *cache.Cache) {
	weatherService := services.NewWeatherService(config, cache)
	WeatherHandler := handlers.NewWeatherHandler(weatherService)

	router := gin.Default()
	router.Use(middleware.LimitByRequest())
	router.GET("/weather/:city", WeatherHandler.GetWeatherByCity)
	router.Run(":" + config.Port)
}
