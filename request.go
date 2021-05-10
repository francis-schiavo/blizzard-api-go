package blizzard_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type BattleNetHeaders struct {
	Namespace string `json:"namespace"`
	Schema    string `json:"schema"`
	Revision  string `json:"revision"`
}

type ApiResponse struct {
	Status           int
	BattleNetHeaders *BattleNetHeaders
	Cached           bool
	Body             []byte
	Error            error
}

func (response ApiResponse) Parse(data interface{}) error {
	return json.Unmarshal(response.Body, data)
}

type ApiClient struct {
	httpClient    *http.Client
	cacheProvider CacheProvider
	game          Game
	region        Region
	token         string
	Classic       bool
	rateLimiter   *RateLimiter
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

func (client ApiClient) Request(url string, query *url.Values, options *RequestOptions) *ApiResponse {
	fullUrl := fmt.Sprintf("%s?%s", url, query.Encode())

	if client.cacheProvider != nil {
		found, cachedResponse := client.cacheProvider.LoadFromCache(fullUrl)
		if found {
			cachedResponse.Cached = true
			return cachedResponse
		}
	}

	request, _ := http.NewRequest(http.MethodGet, fullUrl, nil)
	request.Header.Set("Accept", "application/json")
	if options.Token != "" {
		request.Header.Set("Authorization", "Bearer "+options.Token)
	} else {
		request.Header.Set("Authorization", "Bearer "+client.token)
	}
	token := client.rateLimiter.Acquire()
	response, err := client.httpClient.Do(request)
	client.rateLimiter.Release(token)

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

	apiResponse := &ApiResponse{
		Status: response.StatusCode,
		BattleNetHeaders: &BattleNetHeaders{
			Namespace: response.Header.Get("Battlenet-Namespace"),
			Schema:    response.Header.Get("Battlenet-Schema"),
			Revision:  response.Header.Get("Battlenet-Schema-Revision"),
		},
		Cached: false,
		Body:   bodyData,
		Error:  nil,
	}

	if client.cacheProvider != nil {
		client.cacheProvider.SaveToCache(fullUrl, apiResponse)
		return apiResponse
	}

	return apiResponse
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

func (client *ApiClient) SetAccessToken(token string) {
	client.token = token
}