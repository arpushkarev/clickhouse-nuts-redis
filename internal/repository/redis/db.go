package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	addres   = "localhost:63791"
	password = ""
	DB       = 0
	duration = time.Duration(60 * 1e9)
)

type RedisClient struct {
	*redis.Client
	Duration time.Duration
}

func New() (*RedisClient, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     addres,
		Password: password,
		DB:       DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis client can't connect, error: %w", err)
	}

	return &RedisClient{client, duration}, nil
}
