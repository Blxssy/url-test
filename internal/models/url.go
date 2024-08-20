package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model

	OriginalURL string `gorm:"not null"`
	ShortURL    string `gorm:"uniqueIndex;not null"`
}
