package ratelimiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

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
	Init(name string, limit int64, window time.Duration, storeType StoreType, config RateLimiterConfig)
	Stop()
}

type RateLimiterConfig struct {
	// name     string
	Rdb *redis.Client
	// duration time.Duration
	StopChan chan bool
	Ctx      context.Context
}

type Store interface {
	getStatus() (int64, bool, error)
	incrementAndCheck() (bool, error)
	init(key string, limit int64, window time.Duration, config RateLimiterConfig)
}
