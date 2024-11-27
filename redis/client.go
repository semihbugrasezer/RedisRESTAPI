package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var Client *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Redis not connecting", err)
	}
	log.Println("Redis connect success")
}
