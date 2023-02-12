// redis handles the redis connection
package redis

import (
	"context"
	"os"
	"time"

	gredis "github.com/go-redis/redis/v8"
	"github.com/kustavo/benchmark/go/domain/message"
)

type redis struct {
	client *gredis.Client
}

// NewRedis creates a new redis connection
func NewRedis() (*redis, error) {

	dsn := os.Getenv("REDIS_ADDR")
	client := gredis.NewClient(&gredis.Options{
		Addr: dsn,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, err
	}

	return &redis{client: client}, nil
}

// Set sets a key with a value and a ttl
func (r *redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.client.Set(ctx, key, value, expiration).Err()
	return err
}

// Get returns the value of a key
func (r *redis) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == gredis.Nil {
			err = message.ErrItemNotFound
		}
		return "", err
	}
	return result, nil
}

// Delete deletes a key
func (r *redis) Del(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	return err
}

// Close closes the redis connection
func (r *redis) Close() {
	r.client.Close()
}
