package blizzard_api

import (
	"github.com/go-redis/redis"
	"net/http"
)

type WoWClient struct {
	ApiClient
}

func NewWoWClient(region Region, redisClient *redis.Client, validCacheStatus []int) *WoWClient {
	if validCacheStatus == nil {
		validCacheStatus = []int{200, 404}
	}

	return &WoWClient{
		ApiClient{
			httpClient:       new(http.Client),
			redisClient:      redisClient,
			game:             "wow",
			region:           region,
			validCacheStatus: validCacheStatus,
		},
	}
}

// Achievement API

func (client *WoWClient) AchievementCategoryIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/achievement-category/index", options)
}

func (client *WoWClient) AchievementCategory(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/achievement-category/%d", options, id)
}

func (client *WoWClient) AchievementIndex(options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/achievement/index", options)
}

func (client *WoWClient) Achievement(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(DATA, STATIC_NS,  "/achievement/%d", options, id)
}

func (client *WoWClient) AchievementMedia(id int, options *RequestOptions) *ApiResponse {
	return client.ApiRequest(MEDIA, STATIC_NS,  "/achievement/%d", options, id)
}