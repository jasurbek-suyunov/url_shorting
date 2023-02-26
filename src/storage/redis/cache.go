package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type cache struct {
	rdb *redis.Client
}

func (c *cache) Contains(ctx context.Context, key string) (bool, error) {
	have, err := c.rdb.Exists(ctx, key).Result()
	if have == 1 && err == nil {
		return true, nil
	} else if have == 0 && err == nil {
		return false, nil
	} else {
		return false, err
	}
}

// Delete implements storage.RedisI
func (c *cache) Delete(ctx context.Context, key string) error {
	return c.rdb.Del(ctx, key).Err()
}

// Get implements storage.RedisI
func (c *cache) Get(ctx context.Context, key string) (value string, err error) {

	result := c.rdb.Get(ctx, key)
	if result.Err() != nil {
		return "", result.Err()
	}

	res, err := result.Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

// Set implements storage.RedisI
func (c *cache) Set(ctx context.Context, key string, value string, expTime time.Duration) error {
	fmt.Println("expTime: ", expTime)
	resp := c.rdb.Set(ctx, key, value, expTime*time.Second)
	if resp.Err() != nil {
		return resp.Err()
	}

	return nil
}

func NewCache(rdb *redis.Client) *cache {
	return &cache{rdb}
}
