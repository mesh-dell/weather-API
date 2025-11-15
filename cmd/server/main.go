package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mesh-dell/weather-API/internal/config"
)

func main() {
	config := config.GetConfig()
	router := gin.Default()
	router.Run(":" + config.Port)
}
