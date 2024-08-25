package ratelimiter

import (
	"testing"
	"time"
)

// Mock implementation of Store for testing
type mockStore struct {
	count  int64
	limit  int64
	window time.Duration
}

func (m *mockStore) init(key string, limit int64, window time.Duration, config RateLimiterConfig) {
	m.count = 0
	m.limit = limit
	m.window = window
}

func (m *mockStore) getStatus() (int64, bool, error) {
	return m.count, m.count < m.limit, nil
}

func (m *mockStore) incrementAndCheck() (bool, error) {

	m.count++
	return m.count <= m.limit, nil
}

// Test for initialization
func TestRateLimiter_Init(t *testing.T) {
	rl := &RateLimiter{}
	rl.Init("testLimiter", 10, time.Minute, MEMORY, RateLimiterConfig{})

	if rl.store == nil {
		t.Error("Expected store to be initialized, got nil")
	}

	// Check if stopChan is initialized
	if rl.config.StopChan == nil {
		t.Error("Expected stopChan to be initialized, got nil")
	}
}

// Test for Stop method
func TestRateLimiter_Stop(t *testing.T) {
	rl := &RateLimiter{}
	rl.Init("testLimiter", 10, time.Minute, MEMORY, RateLimiterConfig{})

	// Start in a separate goroutine to avoid blocking
	go rl.Stop()

	select {
	case <-rl.config.StopChan:
		// Stop channel was sent a value, which is expected
	case <-time.After(time.Second):
		t.Error("Expected stopChan to receive a value within 1 second")
	}
}

// Test for GetStatus method
func TestRateLimiter_GetStatus(t *testing.T) {
	rl := &RateLimiter{}
	rl.store = &mockStore{limit: 10}
	rl.Init("testLimiter", 10, time.Minute, MEMORY, RateLimiterConfig{})

	count, ok, err := rl.GetStatus()
	if err != nil {
		t.Errorf("GetStatus() returned error: %v", err)
	}

	if count != 0 {
		t.Errorf("Expected count to be 0, got %d", count)
	}

	if !ok {
		t.Error("Expected status to be OK")
	}
}

// Test for Hit method
func TestRateLimiter_Hit(t *testing.T) {
	rl := &RateLimiter{}
	rl.Init("testLimiter", 2, time.Minute, MEMORY, RateLimiterConfig{})
	rl.store = &mockStore{limit: 2}

	// First hit should be allowed
	ok, err := rl.Hit()
	if err != nil {
		t.Errorf("Hit() returned error: %v", err)
	}

	if !ok {
		t.Error("Expected hit to be allowed")
	}

	// Second hit should be allowed
	ok, err = rl.Hit()
	if err != nil {
		t.Errorf("Hit() returned error: %v", err)
	}

	if !ok {
		t.Error("Expected hit to be allowed")
	}

	// Third hit should be denied
	ok, err = rl.Hit()
	if err != nil {
		t.Errorf("Hit() returned error: %v", err)
	}

	if ok {
		t.Error("Expected hit to be denied")
	}
}
