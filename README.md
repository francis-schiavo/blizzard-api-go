# Blizzard API Go

I created this project because I need to cache responses in
Redis and I wanted to isolate data structs from the actual request.

The reason for that is to have my project-specific structs so I can
benefit from `gopg` tags for my ORM.

If you want to use a more complete library for dealing with Blizzard APIs
I strongly recommend using this one: https://github.com/FuzzyStatic/blizzard

## Usage

```Go
package main

import (
    "fmt"
    blizzard_api "github.com/francis-schiavo/blizzard-api-go"
    "github.com/go-redis/redis"
) 

var RedisClient *redis.Client

type LocalizedField struct {
	EnUS string `json:"en_US"`
	EsMX string `json:"es_MX"`
	PtBR string `json:"pt_BR"`
	DeDE string `json:"de_DE"`
	EnGB string `json:"en_GB"`
	EsES string `json:"es_ES"`
	FrFR string `json:"fr_FR"`
	ItIT string `json:"it_IT"`
	RuRU string `json:"ru_RU"`
	KoKR string `json:"ko_KR"`
	ZhTW string `json:"zh_TW"`
	ZhCN string `json:"zh_CN"`
}

type PowerType struct {
	ID int `json:"id"`
	Name LocalizedField `json:"name"`
}

func main() {
    redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       1,
	})

    wowClient := blizzard_api.NewWoWClient("us", redisClient, nil, false)
    wowClient.CreateAccessToken("YOUR CLIENT ID", "YOUR CLIENT SECRET", "US")
    data := wowClient.PowerType(0, nil)

    var power PowerType
    err := data.Parse(&power)
    if err != nil {
        fmt.Println("Big mistake!")
    }
    fmt.Printf("Power %d is: %s", power.ID, power.Name.EnUS)
}
```