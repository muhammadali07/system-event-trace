package graphql

import (
	"github.com/yourusername/yourapp/internal/service"
)

type Resolver struct {
	service *service.Service
}

func NewResolver(srv *service.Service) *Resolver {
	return &Resolver{service: srv}
}

// Resolver methods implementing GraphQL queries and mutations
