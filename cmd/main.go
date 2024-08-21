package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/Blxssy/url-test/internal/config"
	container2 "github.com/Blxssy/url-test/internal/container"
	"github.com/Blxssy/url-test/internal/logger"
	"github.com/Blxssy/url-test/internal/router"
	"github.com/Blxssy/url-test/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}

	cfg := config.InitConfig()

	log := logger.SetupLogger(cfg.Env)

	log.Info("cfg", slog.Any("cfg", cfg))

	mainStorage := storage.NewStorage(log, cfg)

	container := container2.NewContainer(mainStorage, cfg, log, cfg.Env)

	r := gin.Default()
	router.InitRoutes(r, container)

	r.Run(":8080")
}
