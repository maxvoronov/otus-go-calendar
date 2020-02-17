package main

import "log"

func main() {
	schedulerInstance, err := InitializeScheduler()
	if err != nil {
		log.Fatalf("Scheduler init error: %s", err)
	}

	schedulerInstance.Start()
}
