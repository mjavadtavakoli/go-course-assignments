package main

import (
    "fmt"
    "sync"
    "time"
)

// RateLimiter main structure
type RateLimiter struct {
    tokens      float64
    maxTokens   float64
    refillRate  float64
    lastRefill  time.Time
    mu          sync.Mutex
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(maxTokens float64, refillRate float64) func() bool {
    rl := &RateLimiter{
        tokens:     maxTokens,
        maxTokens:  maxTokens,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }

    // Main closure that returns rate limit status
    return func() bool {
        rl.mu.Lock()
        defer rl.mu.Unlock()

        // Calculate tokens refilled since last check
        now := time.Now()
        elapsed := now.Sub(rl.lastRefill).Seconds()
        rl.tokens += elapsed * rl.refillRate

        // Ensure tokens don't exceed maximum
        if rl.tokens > rl.maxTokens {
            rl.tokens = rl.maxTokens
        }
        rl.lastRefill = now

        // Check token balance
        if rl.tokens >= 1 {
            rl.tokens--
            return true // Request allowed
        }

        return false // Request denied (rate limited)
    }
}

// RateLimiterWithCallback advanced version with callback
func RateLimiterWithCallback(maxTokens float64, refillRate float64, onLimit func()) func() bool {
    rl := &RateLimiter{
        tokens:     maxTokens,
        maxTokens:  maxTokens,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }

    return func() bool {
        rl.mu.Lock()
        defer rl.mu.Unlock()

        now := time.Now()
        elapsed := now.Sub(rl.lastRefill).Seconds()
        rl.tokens += elapsed * rl.refillRate

        if rl.tokens > rl.maxTokens {
            rl.tokens = rl.maxTokens
        }
        rl.lastRefill = now

        if rl.tokens >= 1 {
            rl.tokens--
            return true
        }

        // Execute callback when rate limit is triggered
        if onLimit != nil {
            onLimit()
        }
        return false
    }
}

func main() {
    fmt.Println("Example 1: Simple Rate Limiter")
    
    // Create rate limiter: max 5 tokens, refill 1 token per second
    limiter := NewRateLimiter(5, 1.0)
    
    for i := 1; i <= 10; i++ {
        if limiter() {
            fmt.Printf("Request %d: Allowed\n", i)
        } else {
            fmt.Printf("Request %d: Rate Limited (please wait)\n", i)
        }
        time.Sleep(200 * time.Millisecond)
    }

    fmt.Println("\nExample 2: Rate Limiter with Callback")
    
    // Rate limiter with callback
    limiterWithCallback := RateLimiterWithCallback(3, 0.5, func() {
        fmt.Println("⚠️  Rate limit triggered! Please wait...")
    })

    for i := 1; i <= 15; i++ {
        if limiterWithCallback() {
            fmt.Printf("Request %d: Success\n", i)
        }
        time.Sleep(300 * time.Millisecond)
    }

    fmt.Println("\nExample 3: Multiple Rate Limiters for Different APIs")
    
    // Create different rate limiters with closures
    apiLimiters := map[string]func() bool{
        "api/login":    NewRateLimiter(10, 2), // 10 requests, refill 2 per second
        "api/search":   NewRateLimiter(30, 5), // 30 requests, refill 5 per second
        "api/download": NewRateLimiter(3, 0.5), // 3 requests, refill 0.5 per second
    }

    // Test rate limiters for different APIs
    endpoints := []string{"api/login", "api/search", "api/download", "api/login", "api/download"}

    for _, endpoint := range endpoints {
        if apiLimiters[endpoint]() {
            fmt.Printf("%s: Request allowed\n", endpoint)
        } else {
            fmt.Printf("%s: Request rate limited\n", endpoint)
        }
        time.Sleep(100 * time.Millisecond)
    }
}