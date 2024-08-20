package container

import (
	"github.com/Blxssy/url-test/internal/config"
	"github.com/Blxssy/url-test/internal/storage"
	"log/slog"
)

type Container interface {
	GetStorage() storage.Storage
	GetConfig() *config.Config
	GetLogger() *slog.Logger
	GetEnv() string
	GetUseMemo() bool
}

type container struct {
	storage storage.Storage
	config  *config.Config
	logger  *slog.Logger
	env     string
	useMemo bool
}

func NewContainer(
	storage storage.Storage,
	cfg *config.Config,
	log *slog.Logger,
	env string,
) *container {
	return &container{
		storage: storage,
		config:  cfg,
		logger:  log,
		env:     env,
	}
}

func (c *container) GetStorage() storage.Storage {
	return c.storage
}

func (c *container) GetConfig() *config.Config {
	return c.config
}

func (c *container) GetLogger() *slog.Logger {
	return c.logger
}

func (c *container) GetEnv() string {
	return c.env
}

func (c *container) GetUseMemo() bool {
	return c.config.UseMemo
}
