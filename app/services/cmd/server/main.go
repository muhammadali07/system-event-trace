package main

import (
	"github.com/muhammadali07/service-grap-go-api/app/pkg/config"
	"github.com/muhammadali07/service-grap-go-api/app/pkg/datastore/postgres"
	"github.com/muhammadali07/service-grap-go-api/app/services/internal/api"
)

func main() {
	cfg := config.Load()
	db := postgres.Connect(cfg.Postgres)

	api.Serve(cfg, db)
}
