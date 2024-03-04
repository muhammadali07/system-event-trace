package graphql

import (
	"github.com/muhammadali07/service-grap-go-api/app/services/internal/service"
)

type Resolver struct {
	service *service.Service
}

func NewResolver(srv *service.Service) *Resolver {
	return &Resolver{service: srv}
}

// Resolver methods implementing GraphQL queries and mutations
