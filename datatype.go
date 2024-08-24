package ratelimiter

import "time"

// Define a custom type for the enum
type StoreType int

// Define constants for the enum values
const (
	MEMORY StoreType = iota
	REDIS
)

type RateLimiterInterface interface {
	GetStatus() (int64, bool, error)
	Hit() (bool, error)
	Init(name string, storeType StoreType)
	Stop()
}

type RateLimiterConfig struct {
	name string
	// lazyUpdateSize int64
	// duration       time.Duration
}

type Store interface {
	getStatus() (int64, bool)
	incrementAndCheck() (bool, error)
	init(limit int64, window time.Duration, stopCh *chan bool)
}
