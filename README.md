# weather-API


A lightweight, high-performance weather API built with **Golang**, **Gin**, and **Redis caching**.
It provides fast and consistent weather data retrieval from external providers (e.g., OpenWeatherMap).

---

## 🚀 Features

* **Current weather by city**
* **Redis caching** to reduce external API calls
* **Rate limiting** via middleware
* Clean **service + DTO + config** architecture
* Easy to extend with new providers

---

## 🗂️ Project Structure

```
.
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   └── api.go
│   ├── cache
│   │   └── cache.go
│   ├── config
│   │   └── config.go
│   ├── dtos
│   │   └── weatherDTO.go
│   ├── handlers
│   │   └── weather_handler.go
│   ├── middleware
│   │   └── limiter.go
│   └── services
│       └── weather_service.go
├── README.md
```

---

## Configuration

Environment variables:

| Variable         | Description                                        |
| ---------------- | -------------------------------------------------- |
| `API_KEY`        | Your weather provider API key                      |
| `API_BASE_URL`   | Example: `https://api.openweathermap.org/data/2.5` |
| `REDIS_ADDR`     | Redis address, e.g., `localhost:6379`              |
| `PORT`           | Port                                               |

Example `.env`:

```
API_KEY=your_api_key
API_BASE_URL=[https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline](https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline)
REDIS_ADDR=localhost:6379

```

## 🛠️ Running the Server

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Start Redis

```bash
docker run -p 6379:6379 redis
```

### 3. Run the API

```bash
go run cmd/server/main.go
```

---

## 🔒 Rate Limiting

Implemented using `tollbooth`:

```go
lmt := tollbooth.NewLimiter(1, nil)
```

* 1 request per second per IP
* Returns `429 Too Many Requests` when exceeded

---
https://roadmap.sh/projects/weather-api-wrapper-service
