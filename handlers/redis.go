package handlers

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func GetRedisClient() *redis.Client {
	host := os.Getenv("REMOTE_REDIS_ADDR")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("REMOTE_REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", host, port),
		DB:   0,
	})

	return rdb
}
