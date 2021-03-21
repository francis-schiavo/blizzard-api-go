package blizzard_api

import "time"

type RateLimiterToken time.Time

type RateLimiter struct {
	Window time.Duration
	Concurrency int

	availableToken chan RateLimiterToken
}

func NewRateLimiter(concurrency int, timeWindow time.Duration) *RateLimiter {
	return &RateLimiter{
		Window:         timeWindow,
		Concurrency:    concurrency,
		availableToken: make(chan RateLimiterToken, concurrency),
	}
}

func (limiter *RateLimiter) Acquire() RateLimiterToken {
	return <- limiter.availableToken
}

func (limiter *RateLimiter) Release(token RateLimiterToken)  {
	elapsed := time.Now().Sub(time.Time(token))
	if elapsed < limiter.Window {
		time.Sleep(limiter.Window - elapsed)
	}

	limiter.availableToken <- token
}
