package utils

import (
	"sync"
	"time"
)

// CacheItem represents an item stored in cache
type CacheItem struct {
	Value     interface{}
	ExpiresAt time.Time
	CreatedAt time.Time
}

// IsExpired checks if the cache item has expired
func (ci *CacheItem) IsExpired() bool {
	return time.Now().After(ci.ExpiresAt)
}

// Cache represents an in-memory cache with TTL support
type Cache struct {
	mu              sync.RWMutex
	items           map[string]*CacheItem
	ttl             time.Duration
	maxSize         int
	cleanupInterval time.Duration
	stopChan        chan struct{}
	stopped         bool
}

// NewCache creates a new cache instance
func NewCache(ttl time.Duration, maxSize int) *Cache {
	cache := &Cache{
		items:           make(map[string]*CacheItem),
		ttl:             ttl,
		maxSize:         maxSize,
		cleanupInterval: 5 * time.Minute,
		stopChan:        make(chan struct{}),
		stopped:         false,
	}

	// Start background cleanup goroutine
	go cache.cleanupExpired()

	return cache
}

// Set stores a value in the cache
func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Check if cache is full
	if len(c.items) >= c.maxSize && c.items[key] == nil {
		// Remove the oldest item
		var oldestKey string
		var oldestTime time.Time
		for k, item := range c.items {
			if oldestTime.IsZero() || item.CreatedAt.Before(oldestTime) {
				oldestKey = k
				oldestTime = item.CreatedAt
			}
		}
		if oldestKey != "" {
			delete(c.items, oldestKey)
		}
	}

	c.items[key] = &CacheItem{
		Value:     value,
		ExpiresAt: time.Now().Add(c.ttl),
		CreatedAt: time.Now(),
	}
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	if item.IsExpired() {
		// Don't delete here, let cleanup goroutine handle it
		return nil, false
	}

	return item.Value, true
}

// GetOrSet retrieves a value from cache, or sets it if not found
func (c *Cache) GetOrSet(key string, value interface{}) interface{} {
	if val, ok := c.Get(key); ok {
		return val
	}
	c.Set(key, value)
	return value
}

// Delete removes an item from the cache
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// Clear clears all items from the cache
func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[string]*CacheItem)
}

// Size returns the number of items in the cache
func (c *Cache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

// Exists checks if a key exists in the cache and is not expired
func (c *Cache) Exists(key string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return false
	}

	return !item.IsExpired()
}

// cleanupExpired removes expired items from the cache
func (c *Cache) cleanupExpired() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-c.stopChan:
			// Graceful shutdown
			return
		case <-ticker.C:
			c.mu.Lock()
			for key, item := range c.items {
				if item.IsExpired() {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		}
	}
}

// Stop gracefully stops the cleanup goroutine
func (c *Cache) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.stopped {
		close(c.stopChan)
		c.stopped = true
	}
}

// GetStats returns cache statistics
func (c *Cache) GetStats() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	expiredCount := 0
	for _, item := range c.items {
		if item.IsExpired() {
			expiredCount++
		}
	}

	return map[string]interface{}{
		"total_items":   len(c.items),
		"expired_items": expiredCount,
		"max_size":      c.maxSize,
		"ttl":           c.ttl.String(),
	}
}
