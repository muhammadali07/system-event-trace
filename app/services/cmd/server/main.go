package main

import (
	"github.com/muhammadali07/service-grap-go-api/app/services/internal/api"
	"github.com/muhammadali07/service-grap-go-api/app/services/pkg/config"
	"github.com/muhammadali07/service-grap-go-api/app/services/pkg/datastore/postgres"
)

func main() {
	cfg := config.Load()
	db := postgres.Connect(cfg.Postgres)

	api.Serve(cfg, db)
}
