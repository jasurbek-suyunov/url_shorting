package redis

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/SuyunovJasurbek/url_shorting/config"
	"github.com/SuyunovJasurbek/url_shorting/src/storage"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type redisCache struct {
	rdb *redis.Client

	token storage.TokenI
	user  storage.UserI
}

// consts for redis connection
const (
	readTimeout = 10 * time.Second // 10 seconds
)

func NewRedisCache(cfg config.Config, expires time.Duration) (*redisCache, error) {

	// ...1: creating context
	var ctx context.Context = context.Background()

	val := os.Getenv("REDIS_POOL_SIZE")
	if val == "" {
		return nil, errors.New("REDIS_POOL_SIZE not set")
	}

	// ...2: opening connection to redis
	rdb := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password:    "",          // no password set
		DB:          cfg.RedisDB, // use default DB
		PoolTimeout: readTimeout,
		PoolSize:    cfg.RedisPoolSize,
	})

	// ...3: checking connection
	pong := rdb.Ping(ctx)
	_, err := pong.Result()
	if err != nil {
		return nil, errors.New("cannot connect to redis")
	}

	// ...4: returning redis cache db
	return &redisCache{
		rdb: rdb,
	}, nil
}

// ...1: Token
func (r *redisCache) Token() storage.TokenI {
	if r.token == nil {
		// r.token = NewTokenCacheRepo(r.rdb)
	}

	return r.token
}

// ...2: User
func (r *redisCache) User() storage.UserI {
	if r.user == nil {
		// r.user = NewUserCacheRepo(r.rdb)
	}

	return r.user
}

// ///////////  TODO: BELOW FROM THERE SHOULD BE FIXED
func (r *redisCache) Contains(ctx context.Context, key string) (bool, error) {
	have, err := r.rdb.Exists(ctx, key).Result()
	if have == 1 && err == nil {
		return true, nil
	} else if have == 0 && err == nil {
		return false, nil
	} else {
		return false, err
	}
}

func (r *redisCache) LogOut(ctx context.Context, key string) error {
	return r.rdb.Del(ctx, key).Err()
}
