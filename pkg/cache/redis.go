package cache

import (
	"context"
	"gangbu/pkg/util"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

type redisCache struct {
}

func (r *redisCache) GetString(ctx context.Context, key string) (string, error) {
	client := getRedisClient()
	defer client.Close()
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		// 判断是否是nil错误
		if err == redis.Nil {
			return "", nil
		}
		util.Logger.Errorf("redis get error: %v", err)
		return "", err
	}
	return val, nil
}

func (r *redisCache) SetString(ctx context.Context, key string, value string, expiration time.Duration) error {
	client := getRedisClient()
	defer client.Close()
	err := client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		util.Logger.Errorf("redis set error: %v", err)
		return err
	}
	return nil
}

func (r *redisCache) DeleteString(ctx context.Context, key string) error {
	client := getRedisClient()
	defer client.Close()
	err := client.Del(ctx, key).Err()
	if err != nil {
		util.Logger.Errorf("redis delete error: %v", err)
		return err
	}
	return nil
}

func NewRedisCache() Cache {
	return &redisCache{}
}

func getRedisClient() *redis.Client {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_PASS"))
	client := redis.NewClient(opt)
	return client
}
