package utils

import (
	"sync"
)

var (
	blacklist = make(map[string]struct{})
	mu        sync.Mutex
)

func AddToBlacklist(token string) {
	mu.Lock()
	defer mu.Unlock()
	blacklist[token] = struct{}{}
}

func IsBlacklisted(token string) bool {
	mu.Lock()
	defer mu.Unlock()
	_, exists := blacklist[token]
	return exists
}