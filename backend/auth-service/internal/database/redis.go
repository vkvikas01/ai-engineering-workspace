package database

import (
	"context"
	"ai-engineering-workspace/auth-service/internal/config"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost + ":" + cfg.RedisPort,
	})
}

func PingRedis(rdb *redis.Client) error {
	return rdb.Ping(context.Background()).Err()
}