package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	"service/pkg/models/domain/config"
)

type ConfigManager struct {
	config *config.Config
}

type Config interface {
	GetConfig() *config.Config
}

func NewConfigManager() *ConfigManager {
	configManager := &ConfigManager{}
	configManager.Init()
	return configManager
}

func (cm *ConfigManager) Init() {
	logrus.Debug("Starting config initialization...")
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file: ", err)
		return
	}
	cm.config = &config.Config{
		Database: config.Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
			Driver:   os.Getenv("DB_DRIVER"),
		},
		App: config.App{
			Port: os.Getenv("APP_PORT"),
		},
		Enrich: config.Enrich{
			UrlAge:    os.Getenv("ENRICH_AGE"),
			UrlGender: os.Getenv("ENRICH_GENDER"),
			UrlNation: os.Getenv("ENRICH_NATION"),
		},
	}
	logrus.Info("Config is initialized")
	logrus.Debugf("Full configuration: %+v", cm.config)
}

func (cm *ConfigManager) GetConfig() *config.Config {
	return cm.config
}
