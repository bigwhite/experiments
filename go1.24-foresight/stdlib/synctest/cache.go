package main

import (
	"sync"
	"time"
)

// Cache 是一个泛型并发缓存，支持任意类型的键和值。
type Cache[K comparable, V any] struct {
	mu      sync.Mutex
	items   map[K]cacheItem[V]
	expiry  time.Duration
	creator func(K) V
}

// cacheItem 是缓存中的单个条目，包含值和过期时间。
type cacheItem[V any] struct {
	value     V
	expiresAt time.Time
}

// NewCache 创建一个新的缓存，带有指定的过期时间和创建新条目的函数。
func NewCache[K comparable, V any](expiry time.Duration, f func(K) V) *Cache[K, V] {
	return &Cache[K, V]{
		items:   make(map[K]cacheItem[V]),
		expiry:  expiry,
		creator: f,
	}
}

// Get 返回缓存中指定键的值，如果键不存在或已过期，则创建新条目。
func (c *Cache[K, V]) Get(key K) V {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 检查缓存中是否存在该键
	item, exists := c.items[key]

	// 如果键存在且未过期，返回缓存的值
	if exists && time.Now().Before(item.expiresAt) {
		return item.value
	}

	// 如果键不存在或已过期，创建新条目
	value := c.creator(key)
	c.items[key] = cacheItem[V]{
		value:     value,
		expiresAt: time.Now().Add(c.expiry),
	}

	return value
}
