package ratelimiter

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

const REDIS_STORE_PREFIX = "go-rm-"

type redisStore struct {
	Store
	count    int64
	duration time.Duration
	// sync.RWMutex
	limit int64
	key   string
	rdb   *redis.Client
	ctx   context.Context
}

func (m *redisStore) init(key string, limit int64, window time.Duration, config RateLimiterConfig) {
	m.count = 0
	m.duration = window
	m.limit = limit
	m.key = REDIS_STORE_PREFIX + key
	m.rdb = config.rdb
	if config.ctx != nil {
		m.ctx = config.ctx
	} else {
		m.ctx = context.Background()
	}

}

func (m *redisStore) getStatus() (int64, bool, error) {
	// m.RLock()
	// defer m.RUnlock()
	count, err := m.rdb.Get(m.ctx, m.key).Int64()
	if err != nil {
		return 0, false, err
	} else {
		return count, count < m.limit, nil
	}
}

func (m *redisStore) incrementAndCheck() (bool, error) {
	result, err := m.rdb.Incr(m.ctx, m.key).Result()
	if err != nil {
		return false, err
	}

	if result == 1 {
		// This is the first time the key is set, so set the TTL
		err := m.rdb.Expire(m.ctx, m.key, m.duration).Err()
		if err != nil {
			return false, err
		}
	}

	return result <= int64(m.limit), nil
}
