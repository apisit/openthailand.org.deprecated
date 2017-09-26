package cache

import (
	localRedis "openthailand/cache/redis"
	"os"

	"gopkg.in/redis.v4"
)

type CacheManager struct {
	CacheServiceManager
}

type CacheServiceManager interface {
	Set(key string, value []byte)
	Get(key string) (value []byte, found bool)
}

var (
	REDIS_AWS_URL = os.Getenv("REDIS_AWS_URL")
)

func redisClient() *redis.Client {
	if REDIS_AWS_URL == "" {
		REDIS_AWS_URL = "localhost:6379"
	}
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_AWS_URL,
		Password: "",
		DB:       0,
	})
	return client
}
func NewCacheManager() CacheManager {
	redisClient := redisClient()
	cacheService := &localRedis.CacheService{Client: redisClient}
	return CacheManager{CacheServiceManager: cacheService}
}
