package postgres

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/muhammadali07/service-grap-go-api/app/services/pkg/config"
)

type DB struct {
	*sql.DB
}

func Connect(cfg config.PostgresConfig) *DB {
	db, err := sql.Open("postgres", cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}
