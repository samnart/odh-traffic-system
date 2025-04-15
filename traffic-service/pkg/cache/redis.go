package cache

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdb *redis.Client
	ctx = context.Background()
) 

// Config holds Redis configuration parameters
type Config struct {
	Host		string
	Port		string
	Password	string
	DB			int
}

// DefaultConfig returns a default Redis configuration
func DefaultConfig() *Config {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("REDIS_PORt")
	if port == "" {
		port = "6379"
	}

	password := os.Getenv("REDIS_PASSWORd")

	return &Config{
		Host: 		"redis",
		Port: 		"6379",
		Password: 	password,
		DB:			0,
	}
}

// InitRedis initializes the Redis client with the provided configuration
func InitRedis(config *Config) error {
	if config == nil {
		config = DefaultConfig()
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

	// Verify connection
	_, err := rdb.Ping(ctx).Result()
	return err
}

// Set stores a key-value pair in Redis with a TTL
func Set(key, value string, ttl time.Duration) error {
	return rdb.Set(ctx, key, value, ttl).Err()
}

// Get retrieves a value by key from Redis
func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil		// Key does not exist
	}
	return val, err
}

// Close closes the Redis connection
func Close() error {
	if rdb != nil {
		return rdb.Close()
	}
	return nil
}