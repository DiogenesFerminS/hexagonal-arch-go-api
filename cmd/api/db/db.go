package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	var err error

	ctx := context.Background()

	config, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))

	if err != nil {
		panic("Connection failed...")
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 30 * time.Minute

	DB, err = pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		panic("Connection failed...")
	}

	err = DB.Ping(ctx)

	if err != nil {
		panic("Connection failed...")
	}

	fmt.Println("Connection successfully...")
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL
	)
  `

	_, err := DB.Exec(context.Background(), createUsersTable)

	if err != nil {
		panic("Could not create users table.")
	}

	fmt.Println("Tables created")
}
