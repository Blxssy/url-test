package memory

import (
	"errors"
	"sync"
)

type MemoryStorage struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
	}
}

func (ms *MemoryStorage) Save(originalURL string, shortURL string) error {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	if _, exists := ms.data[originalURL]; exists {
		return errors.New("URL already exists")
	}

	ms.data[shortURL] = originalURL
	return nil
}

func (ms *MemoryStorage) Get(shortURL string) (string, error) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	originalURL, exists := ms.data[shortURL]
	if !exists {
		return "", errors.New("URL not found")
	}

	return originalURL, nil
}
