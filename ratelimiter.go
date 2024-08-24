package ratelimiter

import "time"

type RateLimiter struct {
	RateLimiterInterface
	config   RateLimiterConfig
	store    Store
	stopChan chan bool
}

func getStore(storeType StoreType) Store {
	switch storeType {
	case MEMORY:
		return &memoryStore{}
	default:
		return &memoryStore{}
	}
}

func (r *RateLimiter) Init(name string, limit int64, window time.Duration, storeType StoreType) {
	r.config = RateLimiterConfig{name}
	r.store = getStore(storeType)
	r.stopChan = make(chan bool)
	r.store.init(limit, window, &r.stopChan)
}

func (r *RateLimiter) Stop() {
	r.stopChan <- true
}

func (r *RateLimiter) GetStatus() (int64, bool, error) {
	a, b := r.store.getStatus()
	return a, b, nil
}

func (r *RateLimiter) Hit() (bool, error) {
	return r.store.incrementAndCheck()
}
