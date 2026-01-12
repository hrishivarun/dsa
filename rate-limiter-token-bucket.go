type RateLimiter struct {
    tokens     float64
    lastRefill time.Time
    capacity   float64
    rate       float64
    mu         sync.Mutex
}

func (rl *RateLimiter) Allow() bool {
    // 1. Lock
    // 2. Calculate time passed since lastRefill
    // 3. Add tokens (TimePassed * Rate)
    // 4. Cap at Capacity
    // 5. If tokens >= 1, deduct 1 and return true
    // 6. Else return false

    rl.mu.Lock()
    defer rl.mu.Unlock()
    tokensToAdd := (float64(time.Since(rl.lastRefill))*rl.rate)/float64(time.Second)
    tokensCap := min(rl.tokens+tokensToAdd, rl.capacity)

    
    rl.lastRefill = time.Now()
    rl.tokens = tokensCap
    if(tokensCap>=1) {
        rl.tokens -= 1
        return true
    }
    
    return false
}