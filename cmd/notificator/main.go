package main

import "log"

func main() {
	notificatorInstance, err := InitializeNotificator()
	if err != nil {
		log.Fatalf("Failed to init notificator: %s", err)
	}

	if err := notificatorInstance.Start(); err != nil {
		notificatorInstance.Logger.Errorf(err.Error())
	}
}
