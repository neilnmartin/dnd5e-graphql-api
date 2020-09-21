package repository

import (
	"github.com/neilnmartin/dnd5e-graphql-api/config"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
)

// A datasource that implements a data repository with CRUD for each resource
type datasource struct {
	UserRepo  domain.UserRepo
	RaceRepo  domain.RaceRepo
	ClassRepo domain.ClassRepo
}

// datasourceFactory creates a datasource based on a string input p.
// p will identify the db or other persistence infrastructure used.
func datasourceFactory(p string) datasource {
	// only one implementation so far, default to mongo
	switch p {
	case "mongodb":
		return createMongoDataSource()
	default:
		return createMongoDataSource()
	}
}

// DB is the data source instance
var DB = datasourceFactory(config.Config.Persistence)
