package handlers

import (
    "context"
    "log"
    "myapp/db"
)

func GetUserByID(id int) {
    var firstname string

    query := `SELECT firstname FROM users WHERE id=$1`
    err := db.Pool.QueryRow(context.Background(), query, id).Scan(&firstname)

    if err != nil {
        log.Printf("Error fetching user by ID: %v\n", err)
        return
    }

    log.Printf("User: %s\n", firstname)
}
