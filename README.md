
## Description

The `go-ratelimiter` Go module provides an implementation of rate limiting mechanisms designed for use both in single-machine applications and distributed systems. Rate limiting is crucial for controlling the rate of operations or requests to ensure fair usage, prevent abuse, and maintain performance.

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
TODO



