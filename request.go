package blizzard_api

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ApiResponse struct {
	Status int
	Cached bool
	Body   []byte
	Error  error
}

func (response ApiResponse) Parse(data interface{}) error {
	return json.Unmarshal(response.Body, data)
}

type ApiClient struct {
	httpClient       *http.Client
	redisClient      *redis.Client
	game             Game
	region           Region
	token            string
	validCacheStatus []int
	Classic          bool
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func (client ApiClient) BaseURL(endpointType EndpointType, region Region) string {
	if region == "" {
		region = client.region
	}

	switch endpointType {
	case COMMUNITY:
		return fmt.Sprintf("https://%s.api.blizzard.com/%s", region, client.game)
	case PROFILE:
		return fmt.Sprintf("https://%s.api.blizzard.com/profile/%s", region, client.game)
	case DATA:
		return fmt.Sprintf("https://%s.api.blizzard.com/data/%s", region, client.game)
	case OAUTH:
		return fmt.Sprintf("https://%s.battle.net/oauth/token", region)
	case MEDIA:
		return fmt.Sprintf("https://%s.api.blizzard.com/data/%s/media", region, client.game)
	case SEARCH:
		return fmt.Sprintf("https://%s.api.blizzard.com/data/%s/search", region, client.game)
	default:
		return ""
	}
}

func (client ApiClient) Namespace(namespace Namespace, classic bool) string {
	if classic && namespace == STATIC_NS {
		namespace = CLASSIC_NS
	}

	return fmt.Sprintf(namespace.String(), client.region.String())
}

func (client ApiClient) searchCache(key string) (bool, *ApiResponse) {
	if client.redisClient == nil {
		return false, nil
	}

	data, err := client.redisClient.HGetAll(key).Result()
	if err == nil && len(data) > 0 {
		status, _ := strconv.Atoi(data["s"])
		for _, valid := range client.validCacheStatus {
			if status == valid {
				return true, &ApiResponse{
					Status: status,
					Cached: true,
					Body:   []byte(data["d"]),
					Error:  nil,
				}
			}
		}
	}
	return false, nil
}

func (client ApiClient) saveCache(key string, status int, body []byte) {
	if client.redisClient == nil {
		return
	}

	cacheData := map[string]interface{}{
		"d": body,
		"s": strconv.Itoa(status),
	}

	client.redisClient.HMSet(key, cacheData)
}

func (client ApiClient) Request(url string, query *url.Values, options *RequestOptions) *ApiResponse {
	fullUrl := fmt.Sprintf("%s?%s", url, query.Encode())

	found, cachedResponse := client.searchCache(fullUrl)
	if found {
		return cachedResponse
	}

	request, _ := http.NewRequest(http.MethodGet, fullUrl, nil)
	request.Header.Set("Accept", "application/json")
	if options.Token != "" {
		request.Header.Set("Authorization", "Bearer "+options.Token)
	} else {
		request.Header.Set("Authorization", "Bearer "+client.token)
	}
	response, err := client.httpClient.Do(request)
	if err != nil {
		return &ApiResponse{
			Status: 0,
			Cached: false,
			Body:   nil,
			Error:  err,
		}
	}

	status := response.StatusCode

	bodyData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &ApiResponse{
			Status: status,
			Cached: false,
			Body:   nil,
			Error:  err,
		}
	}

	client.saveCache(fullUrl, status, bodyData)

	return &ApiResponse{
		Status: status,
		Cached: false,
		Body:   bodyData,
		Error:  nil,
	}
}

func (client ApiClient) ApiRequest(endpointType EndpointType, namespace Namespace, uriPattern string, options *RequestOptions, args ...interface{}) *ApiResponse {
	if options == nil {
		options = &RequestOptions{
			Classic: client.Classic,
		}
	}

	fields := &url.Values{}
	if namespace != NONE_NS {
		fields.Add("namespace", client.Namespace(namespace, options.Classic))
	}
	if options.Locale != All {
		fields.Add("locale", options.Locale.String())
	}
	completeUrl := client.BaseURL(endpointType, options.Region) + fmt.Sprintf(uriPattern, args...)
	if options.QueryString != nil {
		for key, value := range options.QueryString {
			fields.Add(key, value)
		}
	}
	return client.Request(completeUrl, fields, options)
}

func (client *ApiClient) CreateAccessToken(clientID string, clientSecret string, region Region) {
	requestURL := client.BaseURL(OAUTH, region)
	body := strings.NewReader("grant_type=client_credentials")

	request, err := http.NewRequest(http.MethodPost, requestURL, body)
	if err != nil {
		log.Fatal(err)
	}
	request.SetBasicAuth(clientID, clientSecret)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	responseData, err := ioutil.ReadAll(response.Body)

	var tokenData Token
	err = json.Unmarshal(responseData, &tokenData)
	if err != nil {
		log.Fatal(err)
	}
	client.token = tokenData.AccessToken
}
