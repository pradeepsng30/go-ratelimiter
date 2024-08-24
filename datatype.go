package ratelimiter

type RateLimiter interface {
	GetStatus() (int64, bool)
	Hit() (bool, error)
}

// type RateLimiterConfig struct {
// 	name struct
// 	store *Store
// 	lazyUpdateSize int64
// 	duration time.Duration
// }

type Store interface {
	getstatus() (int64, bool)
	incrementAndCheck() (bool, error)
	init()
}
