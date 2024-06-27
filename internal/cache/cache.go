package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)




var RedisClient *redis.Client

func InitRedis(addr, password string) {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password,
        DB:       0,
    })

    _, err := RedisClient.Ping(context.Background()).Result()
    if err != nil {
        log.Fatalln("failed to connect to redis")
    }

    log.Println("connected to redis")
}

func CacheVideo(ctx context.Context, key string, value []byte) error {
    return RedisClient.Set(ctx, key, value, 10*time.Second).Err()
}

func GetCachedVideo(ctx context.Context, key string) ([]byte, error) {
    val, err := RedisClient.Get(ctx, key).Bytes()
    if err == redis.Nil {
        return nil, nil // cache miss
    }
    return val, err
}


