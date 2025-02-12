package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"github.com/mngcndl/shortener/internal/storage"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
const shortURLLength = 10
var store storage.Storage

func main() {
	port := getPort()

	if os.Getenv("POSTGRES_CONN") == "postgres" {
		connStr := os.Getenv("POSTGRES_CONN")
		var err error
		store, err = storage.NewPostgresStorage(connStr)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	} else {
		store = storage.NewMemoryStorage()
	}

	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	log.Printf("Starting server on the port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
    originalURL := r.URL.Query().Get("url")
	if originalURL == "" {
		http.Error(w, "Missing URL parameter", http.StatusBadRequest)
		return
	}

	shortURL, err := generateShortURL()
    if err != nil {
        http.Error(w, "Failed to generate short URL", http.StatusInternalServerError)
        return
    }
	
	store.Save(shortURL, originalURL)
    fmt.Fprintf(w, "Shortened URL: %s", shortURL)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	originalURL, found, err := store.Get(shortURL)
	if err != nil {
		http.Error(w, "Failed to retrieve URL", http.StatusInternalServerError)
		return
	}
	if !found {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	// w.WriteHeader(http.StatusNotImplemented)
	// fmt.Fprintln(w, "Redirect API not implemented")
	http.Redirect(w, r, originalURL, http.StatusFound)
}

func generateShortURL() (string, error) {
	b := make([]byte, shortURLLength)
	for i := range b {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[n.Int64()]
	}
	return string(b), nil
}