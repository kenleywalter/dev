package handlers

import (
    "context"
    "log"
    "myapp/db"
)

// BoardGame represents a board game
type BoardGame struct {
    ID     int
    Name   string
    Genre  string
    Year   int
}

// AddBoardGame inserts a new board game into the database
func AddBoardGame(name, genre string, year int) (*BoardGame, error) {
    query := `INSERT INTO boardgames (name, genre, year) 
              VALUES ($1, $2, $3) 
              RETURNING id, name, genre, year`

    var boardGame BoardGame
    err := db.Pool.QueryRow(context.Background(), query, name, genre, year).Scan(&boardGame.ID, &boardGame.Name, &boardGame.Genre, &boardGame.Year)
    if err != nil {
        log.Printf("Error inserting board game: %v\n", err)
        return nil, err
    }

    return &boardGame, nil
}

// GetBoardGames fetches all board games from the database
func GetBoardGames() ([]BoardGame, error) {
    query := `SELECT id, name, genre, year FROM boardgames`

    rows, err := db.Pool.Query(context.Background(), query)
    if err != nil {
        log.Printf("Error fetching board games: %v\n", err)
        return nil, err
    }
    defer rows.Close()

    var boardGames []BoardGame
    for rows.Next() {
        var boardGame BoardGame
        if err := rows.Scan(&boardGame.ID, &boardGame.Name, &boardGame.Genre, &boardGame.Year); err != nil {
            log.Printf("Error scanning board game row: %v\n", err)
            return nil, err
        }
        boardGames = append(boardGames, boardGame)
    }

    if err := rows.Err(); err != nil {
        log.Printf("Error iterating over rows: %v\n", err)
        return nil, err
    }

    return boardGames, nil
}
