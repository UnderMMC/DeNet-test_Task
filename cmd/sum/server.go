package main

import (
	"DeNet-test_Task/internal/app"
	"log"
)

func main() {
	if err := app.New().Run(); err != nil {
		log.Fatalf("Error to start the app: %v", err)
	}
}
