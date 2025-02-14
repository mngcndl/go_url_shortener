package storage

import (
	"testing"
)

func TestMemoryStorage_SaveAndGet(t *testing.T) {
	store := NewMemoryStorage()

	short := "abc123"
	original := "https://example.com"

	err := store.Save(short, original)
	if err != nil {
		t.Fatalf("Failed to save URL: %v", err)
	}

	retrieved, exists, err := store.Get(short)
	if err != nil {
		t.Fatalf("Failed to get URL: %v", err)
	}
	if !exists {
		t.Fatalf("URL not found")
	}
	if retrieved != original {
		t.Fatalf("Expected %s, got %s", original, retrieved)
	}
}

func TestMemoryStorage_GetShortByOriginal(t *testing.T) {
	store := NewMemoryStorage()

	short := "abc123"
	original := "https://example.com"

	err := store.Save(short, original)
	if err != nil {
		t.Fatalf("Failed to save URL: %v", err)
	}

	retrieved, exists, err := store.GetShortByOriginal(original)
	if err != nil {
		t.Fatalf("Failed to get short URL: %v", err)
	}
	if !exists {
		t.Fatalf("Short URL not found")
	}
	if retrieved != short {
		t.Fatalf("Expected %s, got %s", short, retrieved)
	}
}