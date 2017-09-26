package redis

import "gopkg.in/redis.v4"

type CacheService struct {
	Client *redis.Client
}

func (c *CacheService) Set(key string, value []byte) {
	err := c.Client.Set(key, value, 0)
	if err != nil {
		//log.Printf("redis set error %v", err)
	}
	return
}

func (c *CacheService) Get(key string) (value []byte, found bool) {
	cachedValue, err := c.Client.Get(key).Result()
	if err == redis.Nil {
		value = nil
		found = false
	} else if err != nil {
		value = nil
		found = false
		//log.Printf("redis Get error %v", err)
	} else {
		value = []byte(cachedValue)
		found = true
	}
	return
}
