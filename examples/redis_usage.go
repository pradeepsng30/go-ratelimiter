// build: redis_main
package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pradeepsng30/ratelimiter"
)

func main() {
	go replica("1")
	go replica("2")
	go replica("3")
	replica("4")
	time.Sleep(time.Second)
	fmt.Println("stopping now")
}

func replica(name string) {
	// Example using the redis rate limiter
	config := ratelimiter.RateLimiterConfig{
		Rdb: redis.NewClient(&redis.Options{
			Addr: "localhost:6370",
		}),
	}
	rl := &ratelimiter.RateLimiter{}
	rl.Init("test", 25, time.Second, ratelimiter.REDIS, config)

	// Hit the rate limiter and check status
	for i := 0; i < 50; i++ {
		status, err := rl.Hit()
		time.Sleep(time.Millisecond * 50)
		if err != nil {
			fmt.Println(name, "Error:", err)
		} else {
			fmt.Printf("%v | Hit %d: within limit? %v\n", name, i+1, status)
		}
	}
}
