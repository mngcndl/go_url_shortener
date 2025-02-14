package common

import "net/http"

type Storage interface {
    Save(key string, value string) error
    Get(key string) (string, bool, error)
    GetShortByOriginal(original string) (string, bool, error) 
}

type Handler interface {
    CreateShortURL(w http.ResponseWriter, r *http.Request)
    GetOriginalURL(w http.ResponseWriter, r *http.Request)
}

type Service interface {
    CreateShortURL(original string) (string, error)
    GetOriginalURL(short string) (string, error)
}