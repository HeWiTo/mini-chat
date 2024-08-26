package utils

import (
    "github.com/go-redis/redis/v8"
    "context"
)

var Ctx = context.Background()

func NewRedisClientWithHost(host string) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     host + ":6379",
        Password: "",
        DB:       0,
    })

    return rdb
}
