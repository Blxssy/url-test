package service

import (
	"github.com/Blxssy/url-test/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService(t *testing.T) {
	container := test.PrepareForServiceTest()
	service := NewURLService(container)

	tests := []struct {
		name string
		url  string
		size int
	}{
		{
			url:  "http://localhost:8080/12345",
			size: 1,
		},
		{
			url:  "http://localhost:8080/12345",
			size: 5,
		},
		{
			url:  "http://localhost:8080/12345",
			size: 10,
		},
		{
			url:  "http://localhost:8080/12345",
			size: 20,
		},
		{
			url:  "http://localhost:8080/12345",
			size: 30,
		},
	}

	for _, tt := range tests {
		shortURL, err := service.SaveURL(tt.url, tt.size)
		assert.Len(t, shortURL, tt.size)
		assert.NoError(t, err)

		res, err := service.GetOriginalURL(shortURL)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	}
}
