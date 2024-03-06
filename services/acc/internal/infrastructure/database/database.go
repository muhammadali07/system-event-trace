package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Config represents the database configuration.
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

// NewConnection creates a new database connection.
func NewConnection(ctx context.Context, cfg Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("Ping failed:", err)
		return nil, err
	}

	return db, nil
}

// CloseConnection closes the database connection.
func CloseConnection(db *sqlx.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}

// ExampleQuery demonstrates how to execute a query using the database connection.
func ExampleQuery(ctx context.Context, db *sqlx.DB) error {
	// Replace this with your actual query.
	query := `SELECT * FROM accounts`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}

	defer rows.Close()

	// Iterate over the rows and process the results.
	for rows.Next() {
		var (
			id   int
			name string
		)

		err := rows.Scan(&id, &name)
		if err != nil {
			return err
		}

		fmt.Println(id, name)
	}

	return nil
}

func main() {
	// Get the database configuration from environment variables.
	cfg := Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}

	// Create a new database connection.
	db, err := NewConnection(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Close the database connection when the main function exits.
	defer CloseConnection(db)

	// Execute an example query.
	err = ExampleQuery(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
}
