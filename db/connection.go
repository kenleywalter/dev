package db

import (
    "context"
    "fmt"
    "log"
    "time"

    "myapp/config"
    "github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

// Initialize the database connection pool
func InitDB(cfg *config.Config) {
    dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSLMode)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var err error
    Pool, err = pgxpool.New(ctx, dbURL)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }

    // Ensure the connection is valid
    if err := Pool.Ping(ctx); err != nil {
        log.Fatalf("Unable to ping database: %v\n", err)
    }

    log.Println("Database connection established")
}

// CloseDB closes the database connection
func CloseDB() {
    Pool.Close()
}
