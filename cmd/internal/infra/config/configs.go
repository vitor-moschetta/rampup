package config

import (
	"os"

	"github.com/mercadolibre/fury_go-toolkit-config/pkg/config"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
)

func LoadInitialConfig() {
	logger.Info("initializing application configuration")

	if scope := os.Getenv("SCOPE"); scope == "" || scope == "local" {
		setupLocalEnvironment()
	}

	cfg, err := config.Load()
	if err != nil {
		logger.Error("error loading configuration", err)
	}

	ConfigMap[KvsEnabledKey] = cfg.GetString(KvsEnabledKey, "")
	ConfigMap[KvsContainerNameKey] = cfg.GetString(KvsContainerNameKey, "")

	setEnvironments(cfg)
}

func setupLocalEnvironment() {
	_ = os.Setenv("configFileName", "cmd/internal/infra/config/configuration.properties")
	_ = os.Setenv("IS_PROD_SCOPE", "false")
	_ = os.Setenv("checksumEnabled", "false")
}

func setEnvironments(cfg *config.Config) {
	_ = os.Setenv(KvsContainerNameKey, cfg.GetString(KvsContainerNameKey, ""))
	_ = os.Setenv(KvsContainerKey, cfg.GetString(KvsContainerKey, ""))
	_ = os.Setenv(KvsHostReadKey, cfg.GetString(KvsHostReadKey, ""))
	_ = os.Setenv(KvsHostWriteKey, cfg.GetString(KvsHostWriteKey, ""))
	_ = os.Setenv(KvsEnabledKey, cfg.GetString(KvsEnabledKey, ""))
}
