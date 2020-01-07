package config

import (
	"github.com/maxvoronov/otus-go-calendar/internal/domain"
	"github.com/sirupsen/logrus"
)

// Config for application
type Config struct {
	Storage domain.StorageInterface
	Logger  *logrus.Logger
}

// New Init new application config
func New(storage domain.StorageInterface, logger *logrus.Logger) *Config {
	return &Config{
		Storage: storage,
		Logger:  logger,
	}
}
