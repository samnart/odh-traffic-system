package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	Client 	*redis.Client
	Ctx		= context.Background()
)

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:		"localhost:6379",
		Password:	"",
		DB:			0,
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Redis connection failed: %v", err)
	}
	log.Println("Connected to Redis")
}

func Get(key string) (string, error) {
	return Client.Get(Ctx, key).Result()
}

func Set(key string, value string, ttl time.Duration) error {
	return Client.Set(Ctx, key, value, ttl).Err()
}