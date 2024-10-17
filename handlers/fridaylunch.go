package handlers

import (
	"context"
	"log"
	"myapp/db"
	"time"
)

// FridayLunch represents a lunch order
type FridayLunch struct {
	ID   int
	Name string
	Meal string
	Date time.Time
}

// AddFridayLunch inserts a new Friday lunch order
func AddFridayLunch(name, meal string, date time.Time) (*FridayLunch, error) {
	query := `INSERT INTO fridaylunch (name, meal, date) 
              VALUES ($1, $2, $3) 
              RETURNING id, name, meal, date`

	var lunch FridayLunch
	err := db.Pool.QueryRow(context.Background(), query, name, meal, date).Scan(&lunch.ID, &lunch.Name, &lunch.Meal, &lunch.Date)
	if err != nil {
		log.Printf("Error inserting Friday lunch: %v\n", err)
		return nil, err
	}

	return &lunch, nil
}

// GetFridayLunches fetches all Friday lunch orders
func GetFridayLunches() ([]FridayLunch, error) {
	query := `SELECT id, name, meal, date FROM fridaylunch`

	rows, err := db.Pool.Query(context.Background(), query)
	if err != nil {
		log.Printf("Error fetching Friday lunches: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var lunches []FridayLunch
	for rows.Next() {
		var lunch FridayLunch
		if err := rows.Scan(&lunch.ID, &lunch.Name, &lunch.Meal, &lunch.Date); err != nil {
			log.Printf("Error scanning lunch row: %v\n", err)
			return nil, err
		}
		lunches = append(lunches, lunch)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v\n", err)
		return nil, err
	}

	return lunches, nil
}
