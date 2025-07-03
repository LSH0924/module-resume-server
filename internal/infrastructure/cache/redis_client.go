package cache

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() (*redis.Client, error) {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		log.Fatal("REDIS_ADDR 없음")
	}

	opts := &redis.Options{
		Addr: redisAddr,
	}
	rdb := redis.NewClient(opts)

	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	log.Println("Redis connection established")
	return rdb, nil
}
