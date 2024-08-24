package ratelimiter

import (
	"testing"
	"time"
)

// Helper function to create a new memoryStore
func newMemoryStore(limit int64, window time.Duration, stopCh *chan bool) *memoryStore {
	ms := &memoryStore{}
	ms.init(limit, window, stopCh)
	return ms
}

func TestMemoryStoreInitialization(t *testing.T) {
	stopCh := make(chan bool)
	ms := newMemoryStore(5, 10*time.Millisecond*10, &stopCh)

	if ms.limit != 5 {
		t.Errorf("Expected limit to be 5, got %d", ms.limit)
	}

	if ms.duration != 10*time.Millisecond*10 {
		t.Errorf("Expected duration to be 10 seconds, got %v", ms.duration)
	}
}

func TestIncrementAndCheck(t *testing.T) {
	stopCh := make(chan bool)
	ms := newMemoryStore(5, 10*time.Millisecond*10, &stopCh)

	for i := int64(0); i < 5; i++ {
		count, ok := ms.incrementAndCheck()
		if count != i+1 {
			t.Errorf("Expected count to be %d, got %d", i+1, count)
		}
		if !ok {
			t.Errorf("Expected rate to be within limit, but it is not")
		}
	}

	// Check when limit is reached
	count, ok := ms.incrementAndCheck()
	if count != 5 {
		t.Errorf("Expected count to be 5, got %d", count)
	}
	if !ok {
		t.Errorf("Expected rate to be within limit, but it is not")
	}
}

func TestResetAfterWindow(t *testing.T) {
	stopCh := make(chan bool)
	ms := newMemoryStore(5, 1*time.Millisecond*10, &stopCh)

	// Allow time for the ticker to reset the count
	time.Sleep(2 * time.Millisecond * 10)

	// After reset, count should be zero
	count, ok := ms.incrementAndCheck()
	if count != 1 {
		t.Errorf("Expected count to be 1 after reset, got %d", count)
	}
	if !ok {
		t.Errorf("Expected rate to be within limit, but it is not")
	}
}

func TestStopTicker(t *testing.T) {
	stopCh := make(chan bool)
	ms := newMemoryStore(5, 10*time.Millisecond*10, &stopCh)

	// Stop the ticker and ensure it's no longer active
	close(stopCh)
	time.Sleep(1 * time.Millisecond * 10) // Allow time for goroutine to exit

	// Check that incrementAndCheck still functions
	_, ok := ms.incrementAndCheck()
	if ok {
		t.Errorf("Expected rate to be out out to be out of limits")
	}
}
