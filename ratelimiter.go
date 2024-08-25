package ratelimiter

import "time"

type RateLimiter struct {
	RateLimiterInterface
	config RateLimiterConfig
	store  Store
	// stopChan chan bool
}

func getStore(storeType StoreType) Store {
	switch storeType {
	case MEMORY:
		return &memoryStore{}
	case REDIS:
		return &redisStore{}
	default:
		return &memoryStore{}
	}
}

func (r *RateLimiter) Init(name string, limit int64, window time.Duration, storeType StoreType, config RateLimiterConfig) {
	r.config = config
	r.store = getStore(storeType)
	r.config.stopChan = make(chan bool)
	r.store.init(name, limit, window, r.config)
}

func (r *RateLimiter) Stop() {
	r.config.stopChan <- true
}

func (r *RateLimiter) GetStatus() (int64, bool, error) {
	return r.store.getStatus()
}

func (r *RateLimiter) Hit() (bool, error) {
	return r.store.incrementAndCheck()
}
