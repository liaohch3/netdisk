package cache

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

func InitCache() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis init fail: %v", err))
	}
	fmt.Printf("Got ping pong from redis: %v, redis init success\n", pong)
}
