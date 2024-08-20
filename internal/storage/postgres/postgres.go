package postgres

import (
	"github.com/Blxssy/url-test/internal/models"
	"gorm.io/gorm"
)

type PGStorage struct {
	db *gorm.DB
}

func NewPGStorage(db *gorm.DB) *PGStorage {
	return &PGStorage{
		db: db,
	}
}

func (pg *PGStorage) Save(originalURL string, shortURL string) error {
	url := models.URL{
		Model:       gorm.Model{},
		OriginalURL: originalURL,
		ShortURL:    shortURL,
	}
	if err := pg.db.Create(&url).Error; err != nil {
		return err
	}
	return nil
}

func (pg *PGStorage) Get(shortURL string) (string, error) {
	var url models.URL
	if err := pg.db.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		return "", err
	}
	return url.OriginalURL, nil
}
