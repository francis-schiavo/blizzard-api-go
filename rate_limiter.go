package blizzard_api

import "time"

type RateLimiterToken struct {
	AcquiredAt time.Time
}

type RateLimiter struct {
	Window time.Duration
	Concurrency int

	availableTokens chan RateLimiterToken
}

func NewRateLimiter(concurrency int, timeWindow time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		Window:         timeWindow,
		Concurrency:    concurrency,
		availableTokens: make(chan RateLimiterToken, concurrency),
	}

	for t := 1; t <= concurrency; t++ {
		limiter.availableTokens <- RateLimiterToken{
			AcquiredAt: time.Now(),
		}
	}

	return limiter
}

func (limiter *RateLimiter) Acquire() RateLimiterToken {
	token := <- limiter.availableTokens
	token.AcquiredAt = time.Now()
	return token
}

func (limiter *RateLimiter) Release(token RateLimiterToken)  {
	elapsed := time.Now().Sub(token.AcquiredAt)
	if elapsed < limiter.Window {
		time.Sleep(limiter.Window - elapsed)
	}

	limiter.availableTokens <- token
}
