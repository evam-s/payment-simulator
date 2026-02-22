package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var RedisClient redis.Client

var contextBg = context.Background()

func ConnectRedis(url, user, pass string) {
	RedisClient = *redis.NewClient(&redis.Options{
		Addr:     url,
		Username: user,
		Password: pass,
	})

	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis at %s, for user %s: %v", url, user, err)
	} else {
		log.Printf("Connected to Redis at %s, with user %s", url, user)
		if res, _ := RedisClient.Exists(ctx, "poNumber").Result(); res == 0 {
			if res1, _ := RedisClient.Set(ctx, "poNumber", -1, 0).Result(); res1 == "OK" {
				log.Println("PO Number Base Initialised:", res1)
			}
		}
	}
}

func FetchUsingGet(key string) (string, error) {
	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	return RedisClient.Get(ctx, key).Result()
}

func FetchUsingMGet(keys []string) ([]any, error) {
	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	return RedisClient.MGet(ctx, keys...).Result()
}

func FetchUsingHGet(key string) (string, error) {
	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	return RedisClient.Get(ctx, key).Result()
}

func FetchUsingHMGet(key string, fields []string) ([]any, error) {
	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	return RedisClient.HMGet(ctx, key, fields...).Result()
}

func StoreUsingSet(key, value string) (string, error) {
	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	return RedisClient.Set(ctx, key, value, 0).Result()
}

func StoreUsingSetWithTtl(key, value string, ttl time.Duration) (bool, error) {
	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	RedisClient.Set(ctx, key, value, ttl)
	return RedisClient.Expire(ctx, key, ttl).Result()
}

func StoreUsingHSet(key string, values []any) (int64, error) {
	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	return RedisClient.HSet(ctx, key, values...).Result()
}

func StoreUsingHSetWithTtl(key string, values []any, ttl time.Duration) (bool, error) {
	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	RedisClient.HSet(ctx, key, values...)
	return RedisClient.Expire(ctx, key, ttl).Result()
}

func GetNewPoNumber() string {
	var ctx, cancel = context.WithTimeout(contextBg, 5*time.Second)
	defer cancel()
	poNum, _ := RedisClient.Incr(ctx, "poNumber").Result()
	return fmt.Sprintf("%0*d", 10, poNum)
}
