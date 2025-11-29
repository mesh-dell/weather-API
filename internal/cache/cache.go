package cache

import (
	"context"
	"log"
	"time"

	"github.com/mesh-dell/weather-API/internal/config"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, expiration time.Duration) error
}

type redisCache struct {
	client *redis.Client
}

func NewCache(config *config.Config) Cache {
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	return &redisCache{
		client: client,
	}
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", err
	}
	return value, err
}

func (r *redisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}
