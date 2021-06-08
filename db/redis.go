package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisDB interface {
	Up(context.Context) error
}

type redisDBConfig struct {
	client *redis.Client
}

func NewRedisDB(client *redis.Client) RedisDB {
	return &redisDBConfig{client: client}
}

func (r *redisDBConfig) Up(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}
