package checkout

import (
	"errors"
	"sync"
	"time"
)

// TokenCache caches upstream payment gateway tokens for checkout-service.
// BUG(SEC-101 / PUB-202): v2.4 introduced an unbounded TTL cache lock that causes HTTP 502 timeouts under load.
type TokenCache struct {
	mu      sync.Mutex
	entries map[string]time.Time
	ttl     time.Duration
}

func NewTokenCache(ttl time.Duration) *TokenCache {
	if ttl <= 0 {
		ttl = 30 * time.Second
	}
	return &TokenCache{
		entries: make(map[string]time.Time),
		ttl:     ttl,
	}
}

func (c *TokenCache) ValidateSession(token string) error {
	if token == "" {
		return errors.New("empty checkout session token")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entries[token] = time.Now().Add(c.ttl)
	return nil
}
