
## Description
```bash
go get "github.com/pradeepsng30/ratelimiter"
```

The `go-ratelimiter` Go module provides an implementation of rate limiting mechanisms designed for use both in single-machine applications and distributed systems. Rate limiting is crucial for controlling the rate of operations or requests to ensure fair usage, prevent abuse, and maintain performance.

## Installation
```bash
go get "github.com/pradeepsng30/ratelimiter"
```

## Features

### Single-Machine Rate Limiting
- **Memory-Based**: Implements in-memory rate limiting using a simple memory store. Ideal for scenarios where you need to limit rates within a single process or application instance.
- **Time-Based Sliding Windows**: Supports time-based sliding windows to control request rates over a specified duration.
- **Concurrency Safe**: Includes support for concurrent access, ensuring thread safety with appropriate synchronization mechanisms.

### Distributed Rate Limiting
- **External Store Integration**: Designed to work with external storage systems (e.g., Redis, Memcached, or SQL databases) to maintain rate limits across multiple instances of an application.
- **Centralized Control**: Provides a unified rate limiting strategy in distributed environments, leveraging distributed locks or coordination services to enforce limits consistently.
- **Scalable and Fault-Tolerant**: Capable of scaling across multiple machines or containers while handling failover and recovery gracefully.

## Usage

### Single-Machine Rate Limiting
```go
import (
	"fmt"
	"time"

	"github.com/pradeepsng30/ratelimiter"
)

func main() {
	// Example using the in-memory rate limiter
	config := ratelimiter.RateLimiterConfig{}
	rl := &ratelimiter.RateLimiter{}
	rl.Init("test", 10, time.Second, ratelimiter.MEMORY, config)

	// Hit the rate limiter and check status
    status, err := rl.Hit()
    fmt.Printf("Hit within limit? %v\n", status)
}
```

see full code - [here](https://github.com/pradeepsng30/go-ratelimiter/blob/main/examples/memory_usage.go)


### Distributed-System Rate Limiting
```go
import (
	"fmt"
	"time"

	"github.com/pradeepsng30/ratelimiter"
)

func main() {
	// Example using the redis rate limiter
	config := ratelimiter.RateLimiterConfig{
		Rdb: redis.NewClient(&redis.Options{
			Addr: "localhost:6370",
		}),
	}
	rl := &ratelimiter.RateLimiter{}
	rl.Init("test", 10, time.Second, ratelimiter.REDIS, config)

	// Hit the rate limiter and check status
    status, err := rl.Hit()
    fmt.Printf("Hit within limit? %v\n", status)
}
```

see full code - [here](https://github.com/pradeepsng30/go-ratelimiter/blob/main/examples/redis_usage.go)



## Contribution

We welcome contributions to this project! Please raise a issue and Pull Request for the same.

## License

This project is licensed under the [MIT License](LICENSE).
