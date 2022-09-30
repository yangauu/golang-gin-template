package redis

import (
	"log"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

// 连接
func connect() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Println("连接redis失败", err)
	}
}
