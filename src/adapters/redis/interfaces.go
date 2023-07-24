package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type RedisAdapter struct {
	client *redis.Client
	logger *zerolog.Logger
}

func NewRedisAdapter(logger *zerolog.Logger, host, port, password string, db int) RedisAdapter {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	return RedisAdapter{
		client: client,
		logger: logger,
	}
}
