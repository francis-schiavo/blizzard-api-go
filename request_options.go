package blizzard_api

import "time"

type RequestOptions struct {
	Region      Region
	Locale      Locale
	Token       string
	Classic     bool
	Since       *time.Time
	QueryString map[string]string
}
