package storage

import (
	"fmt"
	"github.com/Blxssy/url-test/internal/config"
	"github.com/Blxssy/url-test/internal/models"
	"github.com/Blxssy/url-test/internal/storage/memory"
	pg "github.com/Blxssy/url-test/internal/storage/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

type Storage interface {
	Save(originalURL string, shortURL string) error
	Get(shortURL string) (string, error)
}

func NewStorage(logger *slog.Logger, config *config.Config) Storage {
	if config.UseMemo {
		memo := memory.NewMemoryStorage()
		logger.Info("Successfully create memory storage")
		return memo
	} else {
		db, err := connectDatabase(config)
		if err != nil {
			logger.Error("Failure database connection")
			os.Exit(1)
		}
		logger.Info("Successfully connection to database")
		logger.Info("db", slog.String("port", config.Port))
		db.AutoMigrate(&models.URL{})
		return pg.NewPGStorage(db)
	}
}

func connectDatabase(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.Username,
		config.Database.DBName, config.Database.Password)
	return gorm.Open(postgres.Open(dsn))

}
