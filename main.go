package main

import (
	"github.com/maxvoronov/otus-go-calendar/cmd"
	"github.com/maxvoronov/otus-go-calendar/internal/config"
	"github.com/maxvoronov/otus-go-calendar/storage/sql"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	storage, err := sql.NewStorage(sql.CreateConfigFromEnvironment(), logger)
	if err != nil {
		logger.Fatal(err)
	}

	appConfig := config.New(storage, logger)

	cmd.Execute(appConfig)
}
