package service

import (
	"testing"
	"github.com/mngcndl/go_url_shortener/internal/storage"
)

func TestService_CreateShortURL(t *testing.T) {
	store := storage.NewMemoryStorage()
	svc := NewService(store)

	original := "https://example.com"

	short, err := svc.CreateShortURL(original)
	if err != nil {
		t.Fatalf("Failed to create short URL: %v", err)
	}
	if short == "" {
		t.Fatalf("Expected non-empty short URL, got empty string")
	}

	short2, err := svc.CreateShortURL(original)
	if err != nil {
		t.Fatalf("Failed to create short URL: %v", err)
	}
	if short != short2 {
		t.Fatalf("Expected %s, got %s", short, short2)
	}
}

func TestService_GetOriginalURL(t *testing.T) {
	store := storage.NewMemoryStorage()
	svc := NewService(store)

	original := "https://example.com"
	short, err := svc.CreateShortURL(original)
	if err != nil {
		t.Fatalf("Failed to create short URL: %v", err)
	}

	retrieved, err := svc.GetOriginalURL(short)
	if err != nil {
		t.Fatalf("Failed to get original URL: %v", err)
	}
	if retrieved != original {
		t.Fatalf("Expected %s, got %s", original, retrieved)
	}
}