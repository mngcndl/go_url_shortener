package storage

import "sync"

type MemoryStorage struct {
	mu sync.Mutex
	store map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage {
		store: make(map[string]string),
	}
}

func (s *MemoryStorage) Save(short, original string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[short] = original
	return nil
}

func (s *MemoryStorage) Get(short string) (string, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	original, found := s.store[short]
	if !found {
		return "", false, nil
	}
	return original, true, nil
}

func (s *MemoryStorage) GetShortByOriginal(original string) (string, bool, error) {
	for short, orig := range s.store {
		if orig == original {
			return short, true, nil
		}
	}
	return "", false, nil
}