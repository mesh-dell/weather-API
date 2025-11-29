package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mesh-dell/weather-API/internal/cache"
	"github.com/mesh-dell/weather-API/internal/config"
	dto "github.com/mesh-dell/weather-API/internal/dtos"
)

type WeatherService interface {
	GetWeatherByCity(context context.Context, weatherRequest dto.WeatherRequest) (dto.WeatherResponse, error)
}

type weatherService struct {
	apiKey  string
	apiBase string
	cache   cache.Cache
}

func NewWeatherService(config *config.Config, cache *cache.Cache) WeatherService {
	return &weatherService{
		apiKey:  config.APIKey,
		apiBase: config.APIBase,
		cache:   *cache,
	}
}

func (w *weatherService) GetWeatherByCity(
	context context.Context,
	weatherRequest dto.WeatherRequest,
) (dto.WeatherResponse, error) {
	cacheKey := fmt.Sprintf("weather:%s:%s", weatherRequest.City, weatherRequest.UnitGroup)
	// check cache
	value, err := w.cache.Get(context, cacheKey)
	if err == nil && value != "" {
		var response dto.WeatherResponse
		err := json.Unmarshal([]byte(value), &response)
		if err == nil {
			return response, nil
		}
	}

	// fetch from api
	url := fmt.Sprintf("%s/%s?unitGroup=%s&key=%s&contentType=json",
		w.apiBase,
		weatherRequest.City,
		weatherRequest.UnitGroup,
		w.apiKey,
	)

	res, err := http.Get(url)
	if err != nil {
		return dto.WeatherResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(res.Body)
		return dto.WeatherResponse{}, fmt.Errorf("external API request failed with status %d and error %s",
			res.StatusCode,
			string(respBody),
		)
	}

	var apiResponse struct {
		ResolvedAddress string `json:"resolvedAddress"`
		Days            []struct {
			Temp       float64 `json:"temp"`
			WindSpeed  float64 `json:"windspeed"`
			Conditions string  `json:"conditions"`
		}
	}

	err = json.NewDecoder(res.Body).Decode(&apiResponse)

	if err != nil {
		fmt.Println("Error decoding json response:", err)
		return dto.WeatherResponse{}, err
	}

	weatherResp := dto.WeatherResponse{
		Location:    apiResponse.ResolvedAddress,
		Temperature: apiResponse.Days[0].Temp,
		Windspeed:   apiResponse.Days[0].WindSpeed,
		Conditions:  apiResponse.Days[0].Conditions,
		Time:        time.Now(),
	}

	// cache data
	weatherJSON, _ := json.Marshal(weatherResp)
	err = w.cache.Set(context, cacheKey, string(weatherJSON), 12*time.Hour)
	if err != nil {
		return dto.WeatherResponse{}, err
	}

	return weatherResp, nil
}
