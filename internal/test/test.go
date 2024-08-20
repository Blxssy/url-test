package test

import (
	"github.com/Blxssy/url-test/internal/config"
	"github.com/Blxssy/url-test/internal/container"
	"github.com/Blxssy/url-test/internal/logger"
	"github.com/Blxssy/url-test/internal/storage"
	"log/slog"
)

func PrepareForServiceTest() container.Container {
	conf := createConfig()
	logger := initTestLogger()
	container := initContainer(conf, logger)

	return container
}

func createConfig() *config.Config {
	conf := &config.Config{}
	conf.Database.Dialect = "postgres"
	conf.Database.Host = "localhost"
	conf.Database.DBName = "test"
	conf.Database.Port = "5432"
	conf.Database.Username = "postgres"
	conf.Database.Password = "postgres"
	conf.UseMemo = true
	return conf
}

func initContainer(conf *config.Config, logger *slog.Logger) container.Container {
	rep := storage.NewStorage(logger, conf)
	container := container.NewContainer(rep, conf, logger, "test")
	return container
}

func initTestLogger() *slog.Logger {
	logger := logger.SetupLogger("local")
	return logger
}
