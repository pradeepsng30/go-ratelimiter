// build: memory_main
package main

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
	for i := 0; i < 50; i++ {
		status, err := rl.Hit()
		time.Sleep(time.Millisecond * 50)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Hit %d: within limit? %v\n", i+1, status)
		}
	}
}
