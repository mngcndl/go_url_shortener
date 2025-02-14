package storage

import (
	"testing"
)

func TestPostgresStorage_SaveAndGet(t *testing.T) {
	connStr := "postgres://user:password@localhost/dbname?sslmode=disable"
	store, err := NewPostgresStorage(connStr)
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	_, err = store.db.Exec("DELETE FROM urls")
    if err != nil {
        t.Fatalf("Failed to clear the database table: %v", err)
    }

	short := "jdfhjdhfjhf"
	original := "https://mynewexample.com"

	err = store.Save(short, original)
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
