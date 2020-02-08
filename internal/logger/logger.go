package logger

import "github.com/sirupsen/logrus"

// NewLogger create new logrus instance
func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger
}
