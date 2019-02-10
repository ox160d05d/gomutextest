package main

import (
	"sync"
)

type MemoryCacheInterface interface {
	Get(string) (string, bool)
	Set(string, string)
	Len() uint32
}

type memoryCache struct {
	mu   sync.Mutex
	data map[string]string
}

func NewMemoryCache() MemoryCacheInterface {
	return &memoryCache{data: make(map[string]string)}
}

func (s *memoryCache) Get(key string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if val, ok := s.data[key]; ok {
		return val, true
	} else {
		return "", false
	}
}

func (s *memoryCache) Set(key string, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *memoryCache) Len() uint32 {
	s.mu.Lock()
	defer s.mu.Unlock()

	return uint32(len(s.data))
}
