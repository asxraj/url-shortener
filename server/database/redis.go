package database

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDDIT_PSWD"),
		DB:       dbNo,
	})

	ping, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Error init redis: %v", err)
	}

	log.Println("Redis started successfully:", ping)

	return rdb
}
