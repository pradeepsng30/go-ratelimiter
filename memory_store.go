package ratelimiter

import (
	"time"
)

type memoryStore struct {
	Store
	count    int64
	duration time.Duration
	// sync.RWMutex
	limit int64
}

func (m *memoryStore) init(key string, limit int64, window time.Duration, config RateLimiterConfig) {
	m.count = 0
	m.duration = window
	m.limit = limit
	var stop chan bool
	if config.StopChan != nil {
		stop = config.StopChan
	} else {
		stop = make(chan bool)
	}
	go func() {
		ticker := time.NewTicker(m.duration)
		for {
			select {
			case <-ticker.C:
				// m.Lock()
				m.count = 0
				// m.Unlock()
			case <-stop:
				m.limit = -1
				return
			}
		}
	}()
}

func (m *memoryStore) getStatus() (int64, bool, error) {
	// m.RLock()
	// defer m.RUnlock()
	return m.count, m.count <= m.limit, nil
}

func (m *memoryStore) incrementAndCheck() (bool, error) {
	// m.Lock()
	// defer m.Unlock()
	m.count++
	return m.count <= m.limit, nil
}
