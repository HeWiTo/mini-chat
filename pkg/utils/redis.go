package utils

import (
    "github.com/go-redis/redis/v8"
    "context"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",              
        DB:       0,               
    })

    return rdb
}
