package core

import (
	"sync"
	"time"

	"github.com/ODIN7h3C0d3r/Netra/internal/formatter"
)

// CacheEntry represents a cached IP lookup result
type CacheEntry struct {
	IPInfo   *formatter.IPInfo
	Expiry   time.Time
	Attempts int
}

// IPInfoCache provides thread-safe in-memory caching
type IPInfoCache struct {
	cache map[string]*CacheEntry
	ttl   time.Duration
	mutex sync.RWMutex
}

// NewIPInfoCache creates a new cache with specified TTL
func NewIPInfoCache(ttl time.Duration) *IPInfoCache {
	return &IPInfoCache{
		cache: make(map[string]*CacheEntry),
		ttl:   ttl,
	}
}

// Get retrieves IPInfo from cache if it exists and is not expired
func (c *IPInfoCache) Get(ip string) (*formatter.IPInfo, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, found := c.cache[ip]
	if !found {
		return nil, false
	}

	if time.Now().After(entry.Expiry) {
		return nil, false
	}

	return entry.IPInfo, true
}

// Set stores IPInfo in cache with current timestamp and TTL
func (c *IPInfoCache) Set(ip string, info *formatter.IPInfo) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[ip] = &CacheEntry{
		IPInfo:   info,
		Expiry:   time.Now().Add(c.ttl),
		Attempts: 0,
	}
}

// RecordAttempt records failed attempts to prevent cache stampedes
func (c *IPInfoCache) RecordAttempt(ip string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if entry, exists := c.cache[ip]; exists {
		entry.Attempts++
	} else {
		c.cache[ip] = &CacheEntry{Attempts: 1}
	}
}

// AttemptCount returns the number of failed attempts for an IP
func (c *IPInfoCache) AttemptCount(ip string) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if entry, exists := c.cache[ip]; exists {
		return entry.Attempts
	}
	return 0
}

// Sweep removes expired entries from cache (should be called periodically)
func (c *IPInfoCache) Sweep() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	now := time.Now()
	for ip, entry := range c.cache {
		if now.After(entry.Expiry) {
			delete(c.cache, ip)
		}
	}
}
