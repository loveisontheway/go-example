package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

// Create a new Redis client
func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Redis server address
	})
}

// Acquire a distributed lock
func acquireLock(client *redis.Client, lockKey string, value string, expiration time.Duration) (bool, error) {
	// Try to set the lock
	ok, err := client.SetNX(ctx, lockKey, value, expiration).Result()
	if err != nil {
		return false, err
	}
	return ok, nil
}

// Release a distributed lock
func releaseLock(client *redis.Client, lockKey string, value string) (bool, error) {
	// Use Lua script to ensure the lock is released only if the value matches
	luaScript := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
	`
	result, err := client.Eval(ctx, luaScript, []string{lockKey}, value).Result()
	if err != nil {
		return false, err
	}
	return result.(int64) == 1, nil
}

func main() {
	client := newRedisClient()
	defer client.Close()

	lockKey := "myLock"
	lockValue := "unique_lock_value" // This should be unique for each client
	expiration := 10 * time.Second

	// Acquire the lock
	acquired, err := acquireLock(client, lockKey, lockValue, expiration)
	if err != nil {
		fmt.Println("Error acquiring lock:", err)
		return
	}

	if acquired {
		fmt.Println("Lock acquired!")

		// Simulate some work
		time.Sleep(5 * time.Second)

		// Release the lock
		released, err := releaseLock(client, lockKey, lockValue)
		if err != nil {
			fmt.Println("Error releasing lock:", err)
			return
		}

		if released {
			fmt.Println("Lock released!")
		} else {
			fmt.Println("Failed to release lock: lock value does not match.")
		}
	} else {
		fmt.Println("Failed to acquire lock, it might be held by another process.")
	}
}
