package service

import (
	"github.com/Blxssy/url-test/internal/container"
	"github.com/Blxssy/url-test/internal/utils/random"
)

type URLService interface {
	SaveURL(originalURL string, size int) (string, error)
	GetOriginalURL(shortUrl string) (string, error)
}

type service struct {
	container container.Container
}

func NewURLService(container container.Container) URLService {
	return &service{
		container: container,
	}
}

func (s *service) SaveURL(originalURL string, size int) (string, error) {
	shortURL := random.NewRandomString(size)

	if err := s.container.GetStorage().Save(originalURL, shortURL); err != nil {
		return "", err
	}

	return shortURL, nil
}

func (s *service) GetOriginalURL(shortUrl string) (string, error) {
	originalURL, err := s.container.GetStorage().Get(shortUrl)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
