package utils

import (
	"sync"
)

// In-memory blacklist
var (
	blacklist = make(map[string]struct{})
	mu        sync.Mutex
)

// AddToBlacklist adds a token to the blacklist
func AddToBlacklist(token string) {
	mu.Lock()
	defer mu.Unlock()
	blacklist[token] = struct{}{}
}

// IsBlacklisted checks if a token is in the blacklist
func IsBlacklisted(token string) bool {
	mu.Lock()
	defer mu.Unlock()
	_, exists := blacklist[token]
	return exists
}