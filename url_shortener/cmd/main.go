package main

import (
	"log"
	"net/http"
	"os"
	"github.com/mngcndl/shortener/internal/common"
	"github.com/mngcndl/shortener/internal/storage"
	"github.com/mngcndl/shortener/internal/handler"
	"github.com/mngcndl/shortener/internal/service"
	"github.com/mngcndl/shortener/config"
)

var store common.Storage

func main() {
    cfg := config.LoadConfig()
    var err error
    
    switch cfg.StorageType {
    case "memory":
        store = storage.NewMemoryStorage()
    case "postgres":
        store, err = storage.NewPostgresStorage(cfg.PostgresURL)
        if err != nil {
            log.Fatal("Failed to connect to PostgreSQL: %v", err)
        }
    default:
        log.Fatal("Invalid storage type")
    }
    
    svc := service.NewService(store)
    h := handler.NewHandler(svc)
    
    router := http.NewServeMux()
    router.HandleFunc("/shorten", h.CreateShortURL)
    router.HandleFunc("/{short}", h.GetOriginalURL)
    
    port := getPort()
    log.Printf("Starting server on port %s", port)
    if err := http.ListenAndServe(":"+port, router); err != nil {
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

// func redirectHandler(w http.ResponseWriter, r *http.Request) {
// 	shortURL := r.URL.Path[1:]
// 	originalURL, found, err := store.Get(shortURL)
// 	if err != nil {
// 		http.Error(w, "Failed to retrieve URL", http.StatusInternalServerError)
// 		return
// 	}
// 	if !found {
// 		http.Error(w, "Short URL not found", http.StatusNotFound)
// 		return
// 	}
// 	http.Redirect(w, r, originalURL, http.StatusFound)
// }
