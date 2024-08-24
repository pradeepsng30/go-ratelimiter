package ratelimiter

import (
	"time"
)

type memoryStore struct {
	RateLimiter
	count    int64
	duration time.Duration
	// sync.RWMutex
	limit int64
}

func (m *memoryStore) init(limit int64, window time.Duration, stopCh *chan bool) {
	m.count = 0
	m.duration = window
	m.limit = limit
	var stop chan bool
	if stopCh != nil {
		stop = *stopCh
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

func (m *memoryStore) getstatus() (int64, bool) {
	// m.RLock()
	// defer m.RUnlock()
	return m.count, m.count <= m.limit
}

func (m *memoryStore) incrementAndCheck() (int64, bool) {
	// m.Lock()
	// defer m.Unlock()
	if m.count < m.limit {
		m.count++
	}
	return m.count, m.count <= m.limit
}
