package dto

import "time"

type WeatherResponse struct {
	Location    string    `json:"location"`
	Temperature float64   `json:"temperature"`
	Windspeed   float64   `json:"windspeed"`
	Conditions  string    `json:"conditions"`
	Time        time.Time `json:"time"`
}

type WeatherRequest struct {
	City      string `json:"city"`
	UnitGroup string `json:"unitgroup"`
}
