package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

// DB holds the database connection pool.
var DB *pgx.Conn

func Connect(ctx context.Context) {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
    DB, err = pgx.Connect(ctx, connStr)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
}

func Close(ctx context.Context) {
    if DB != nil {
        if err := DB.Close(ctx); err != nil {
            log.Printf("Error closing database: %v\n", err)
        }
    }
}

func GetGreeting(ctx context.Context) (string, error) {
    var greeting string
    err := DB.QueryRow(ctx, "SELECT 'Hello, world!'").Scan(&greeting)
    if err != nil {
        return "", fmt.Errorf("QueryRow failed: %v", err)
    }
    return greeting, nil
}
