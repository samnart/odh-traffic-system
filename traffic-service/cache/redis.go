package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var rdb *redis.Client
var ctx = context.Background()

func InitRedis() error {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	return err
}

func Set(key, value string, ttl time.Duration) error {
	return rdb.Set(ctx, key, value, ttl).Err()
}

func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}