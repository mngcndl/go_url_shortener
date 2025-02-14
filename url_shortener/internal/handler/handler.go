package handler

import (
	"net/http"
	"github.com/mngcndl/go_url_shortener/internal/common"
	"log"
	"encoding/json"
)

// type Handler struct {
// 	service *service.Service
// }

type handler struct {
	service common.Service
}

func NewHandler(service common.Service) common.Handler {
    return &handler{
        service: service,
    }
}

func (h *handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

	log.Printf("Received POST request to shorten URL: %s", r.URL.Query().Get("url"))
	original := r.FormValue("url")
	if original == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	short, err := h.service.CreateShortURL(original)
	if err != nil {
		log.Printf("Failed to create short URL: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Created short URL: %s -> %s", short, original)
	
	response := map[string]string{
		"original": original,
		"short":    short,
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK) 
    json.NewEncoder(w).Encode(response)
}

func (h *handler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	// short := r.URL.Path[len("/"):]
	short := r.URL.Path[1:]
	original, err := h.service.GetOriginalURL(short)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Received GET request for short URL: %s", short)
	log.Printf("Redirecting short URL: %s -> %s", short, original)
	http.Redirect(w, r, original, http.StatusFound)
}
