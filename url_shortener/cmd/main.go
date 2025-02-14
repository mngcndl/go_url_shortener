package main

import (
    "flag"
	"log"
	"net/http"
	"os"
    "database/sql"
    "time"
    _ "github.com/lib/pq"
	"github.com/mngcndl/shortener/internal/common"
	"github.com/mngcndl/shortener/internal/storage"
	"github.com/mngcndl/shortener/internal/handler"
	"github.com/mngcndl/shortener/internal/service"
	"github.com/mngcndl/shortener/config"
)

var store common.Storage

func main() {
    storageType := flag.String("storage", "postgres", "Storage type (postgres / memory)")
    postgresURL := flag.String("postgres-url", "postgres://user:password@localhost/dbname?sslmode=disable", "URL for accessing PostgreSQL")
    flag.Parse()

    if *storageType == "postgres" && *postgresURL == "" {
        log.Fatal("Flag -postgres-url is required when using postgres storage")
    }
    cfg := config.LoadConfig(*storageType, *postgresURL)
    var err error

    log.Printf("The storage type: %s", cfg.StorageType)
    switch cfg.StorageType {
    case "memory":
        store = storage.NewMemoryStorage()
    case "postgres":
        waitForPostgres(cfg.PostgresURL)
        store, err = storage.NewPostgresStorage(cfg.PostgresURL)
        if err != nil {
            log.Fatalf("Failed to connect to PostgreSQL: %v", err)
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

func waitForPostgres(postgresURL string) {
    for {
        db, err := sql.Open("postgres", postgresURL)
        if err != nil {
            log.Printf("Error opening database: %v", err)
            time.Sleep(2 * time.Second)
            continue
        }

        err = db.Ping()
        if err == nil {
            log.Println("PostgreSQL is ready!")
            return
        }

        log.Println("Waiting for PostgreSQL to be ready...")
        time.Sleep(2 * time.Second)
    }
}