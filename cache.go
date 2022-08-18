package blizzard_api

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

type CacheProvider interface {
	SaveToCache(identifier string, apiResponse *ApiResponse)
	LoadFromCache(identifier string) (bool, *ApiResponse)
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(redisClient *redis.Client) *RedisCache {
	return &RedisCache{client: redisClient}
}

func (cache *RedisCache) LoadFromCache(identifier string) (bool, *ApiResponse) {
	data, err := cache.client.Get(identifier).Bytes()

	if err != nil {
		return false, nil
	}

	var response *ApiResponse
	if json.Unmarshal(data, &response) != nil {
		return false, nil
	}

	return true, response
}

func (cache *RedisCache) SaveToCache(identifier string, apiResponse *ApiResponse) {
	data, _ := json.Marshal(apiResponse)
	cache.client.Set(identifier, data, 0)
}
