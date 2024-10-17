package main

import (
	"log"
	"myapp/config"
	"myapp/db"
	"myapp/handlers"
	"time"
)

func main() {
	// Load configuration from .env
	cfg := config.LoadConfig()

	// Initialize the database connection
	db.InitDB(cfg)
	defer db.Pool.Close()

	// Example: Add a new board game
	newGame, err := handlers.AddBoardGame("Catan", "Strategy", 1995)
	if err != nil {
		log.Fatalf("Failed to add board game: %v", err)
	}
	log.Printf("Created board game: %+v", newGame)

	boardGames, err := handlers.GetBoardGames()
	if err != nil {
		log.Fatalf("Failed to fetch board games: %v", err)
	}
	log.Printf("Fetched board games: %+v", boardGames)

	lunchDate := time.Now().AddDate(0, 0, 5)
	newLunch, err := handlers.AddFridayLunch("Alice", "Pizza", lunchDate)
	if err != nil {
		log.Fatalf("Failed to add Friday lunch: %v", err)
	}
	log.Printf("Created lunch order: %+v", newLunch)

	lunches, err := handlers.GetFridayLunches()
	if err != nil {
		log.Fatalf("Failed to fetch Friday lunches: %v", err)
	}
	log.Printf("Fetched Friday lunches: %+v", lunches)
}
