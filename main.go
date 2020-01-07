package main

import (
	"github.com/maxvoronov/otus-go-calendar/cmd"
	"github.com/maxvoronov/otus-go-calendar/internal/config"
	"github.com/maxvoronov/otus-go-calendar/storage/inmemory"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	appConfig := config.New(inmemory.NewInMemoryStorage(), logger)

	cmd.Execute(appConfig)
}
