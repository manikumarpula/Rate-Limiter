package ratelimiter

import "time"

// RateLimiter defines the interface for rate limiting implementations
type RateLimiter interface {
	// Allow checks if a request should be allowed based on the key
	Allow(key string) bool
	
	// Reset resets the rate limiter for a specific key
	Reset(key string) error
}

// Config represents the configuration for a rate limiter
type Config struct {
	RequestsPerSecond int           `yaml:"requests_per_second"`
	BurstSize         int           `yaml:"burst_size"`
	WindowSize        time.Duration `yaml:"window_size"`
	Algorithm         string        `yaml:"algorithm"`
}

// Response represents the response from a rate limiter check
type Response struct {
	Allowed       bool          `json:"allowed"`
	RemainingRequests int       `json:"remaining_requests"`
	ResetTime     time.Time     `json:"reset_time"`
	RetryAfter    time.Duration `json:"retry_after,omitempty"`
}