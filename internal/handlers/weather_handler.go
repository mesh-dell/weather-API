package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/mesh-dell/weather-API/internal/dtos"
	"github.com/mesh-dell/weather-API/internal/services"
)

type WeatherHandler struct {
	weatherService services.WeatherService
}

func NewWeatherHandler(weatherService services.WeatherService) *WeatherHandler {
	return &WeatherHandler{
		weatherService: weatherService,
	}
}

func (h *WeatherHandler) GetWeatherByCity(c *gin.Context) {
	var weatherRequest dto.WeatherRequest
	weatherRequest.City = c.Param("city")
	weatherRequest.UnitGroup = c.Query("unitGroup")

	if weatherRequest.UnitGroup == "" {
		weatherRequest.UnitGroup = "metric"
	}

	if weatherRequest.City == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "missing city parameter",
		})
		return
	}

	weather, err := h.weatherService.GetWeatherByCity(c.Request.Context(), weatherRequest)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, weather)
}
