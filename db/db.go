package db

import (
	"context"

	"github.com/himalczyk/simple-web-server/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

// DB holds the database connection pool.
type Client struct {
    *pgxpool.Pool
}

func NewClient(ctx context.Context, url string) (*Client, error) {
    poolConfig, err := pgxpool.ParseConfig(url)
    if err != nil {
        return nil, err
    }
    pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
    if err != nil {
        return nil, err
    }
    return &Client{pool}, nil
}

func (c *Client) Close() {
    c.Pool.Close()
}


func (c *Client) RegisterUser(ctx context.Context, userData *models.RegisterData) error {
    _, err := c.Exec(ctx, "INSERT INTO users (username, password, email, favorite_pokemon) VALUES ($1, $2, $3, $4)",
        userData.Username, userData.Password, userData.Email, userData.FavoritePokemon)
    return err
}
