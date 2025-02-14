package service

import (
	"errors"
	"github.com/mngcndl/shortener/internal/common"
	"github.com/mngcndl/shortener/pkg/shortener"
)

type service struct {
	storage common.Storage
}

func NewService(storage common.Storage) common.Service {
	return &service{
		storage: storage,
	}
}

func (s *service) CreateShortURL(original string) (string, error) {
	short, exists, err := s.storage.GetShortByOriginal(original)
	if err != nil {
		return "", err
	}
	if exists {
		return short, nil
	}
	
	short, errSh := shortenHandler.GenerateShortURL()
	if errSh != nil {
		return "", errSh
	}
	if err := s.storage.Save(short, original); err != nil {
		return "", err
	}
	return short, nil
}

func (s *service) GetOriginalURL(short string) (string, error) {
	original, exists, err := s.storage.Get(short)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.New("Short URL not found")
	}
	return original, nil
}