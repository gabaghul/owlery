package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

func (a RedisAdapter) Store(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if err := a.client.Set(ctx, key, value, ttl).Err(); err != nil {
		return errors.Wrap(err, fmt.Sprintf("could not store value for key %s", key))
	}
	return nil
}

func (a RedisAdapter) Get(ctx context.Context, key string) (value interface{}, err error) {
	value, err = a.client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not get value for key %s", key))
	}

	return value, nil
}
