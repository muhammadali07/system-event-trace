package api

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/muhammadali07/service-grap-go-api/app/pkg/config"
	"github.com/muhammadali07/service-grap-go-api/app/services/internal/handler"
	"github.com/muhammadali07/service-grap-go-api/app/services/internal/service"
)

func Serve(cfg *config.Config, db *postgres.DB) {
	srv := service.NewService(db)
	h := handler.NewHandler(srv)

	http.Handle("/graphql", handler.GraphQL(handler.NewExecutableSchema(handler.Config{Resolvers: &handler.Resolver{}})))
	http.Handle("/query", handler.Playground("GraphQL playground", "/graphql"))

	log.Printf("Server running on port %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, nil))
}
