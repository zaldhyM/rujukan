package cache

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

// Connect initializes the Redis client and verifies the connection.
func Connect() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}

	RDB = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := RDB.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("❌ Failed to connect to Redis at %s: %v", addr, err)
	}

	log.Printf("✅ Redis connected at %s", addr)
	return RDB
}

// Close closes the Redis connection.
func Close() {
	if RDB != nil {
		if err := RDB.Close(); err != nil {
			log.Printf("⚠️ Error closing Redis connection: %v", err)
		}
	}
}

// SessionKey returns the Redis key for a user session.
func SessionKey(userID uint) string {
	return fmt.Sprintf("session:%d", userID)
}
