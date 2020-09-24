package graph

import (
	"github.com/neilnmartin/dnd5e-graphql-api/graph/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var db = repository.DB

// Resolver struct holds general resolvers
type Resolver struct {
}
