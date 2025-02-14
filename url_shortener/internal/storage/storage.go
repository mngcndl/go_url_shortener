package storage

import (
    "github.com/mngcndl/shortener/internal/common"
)

type storage struct {}
// type Storage interface {
// 	Save(short, original string) error
// 	Get(short string) (string, bool, error)
// }

func (s *storage) Get(short string) (string, bool, error) {
    // Реализация получения URL
    return "", false, nil
}

func (s *storage) Save(short, original string) error {
    // Реализация сохранения URL
    return nil
}
func NewStorage() common.Storage {
	return &storage{}
}
// func NewStorage(config Config) (Storage, error) {
//     switch config.StorageType {
//     case "memory":
//         return memory.NewMemoryStorage(), nil
//     case "postgres":
//         return postgres.NewPostgresStorage(), nil
//     default:
//         return nil, fmt.Errorf("invalid storage type: %s", config.StorageType)
//     }
// }