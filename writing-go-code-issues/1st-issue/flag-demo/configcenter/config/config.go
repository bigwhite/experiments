package config

import (
	"log"
	"sync"
)

var (
	m  map[string]interface{}
	mu sync.RWMutex
)

func init() {
	m = make(map[string]interface{}, 10)
}

func SetString(k, v string) {
	mu.Lock()
	m[k] = v
	mu.Unlock()
}

func SetInt(k string, i int) {
	mu.Lock()
	m[k] = i
	mu.Unlock()
}

func GetString(key string) string {
	mu.RLock()
	defer mu.RUnlock()
	v, ok := m[key]
	if !ok {
		return ""
	}
	return v.(string)
}

func GetInt(key string) int {
	mu.RLock()
	defer mu.RUnlock()
	v, ok := m[key]
	if !ok {
		return 0
	}
	return v.(int)
}

func Dump() {
	log.Println(m)
}
