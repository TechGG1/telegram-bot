package main

import (
	"log"
	"telegram-bot/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to connect with BeerAdvicer: %s", err)
	}
}
