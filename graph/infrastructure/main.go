package infrastructure

import (
	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
)

// A Datasource that implements a data repository with CRUD for each resource
type Datasource interface {
	CreateUser(domain.User) domain.User
	// ReadUser(domain.User) domain.User
	// UpdateUser(domain.User) domain.User
	// DeleteUser(domain.User) domain.User

	CreateGame()
	ReadGame()
	UpdateGame()
	DeleteGame()

	CreateCharacter()
	ReadCharacter()
	UpdateCharacter()
	DeleteCharacter()
}

// DatasourceFactory creates a Datasource interface
func DatasourceFactory(p string) Datasource {
	switch p {
	// only one implementation so far
	case "mongodb":
		return createMongoDataSource()
	// default to mongo
	default:
		return createMongoDataSource()
	}
}
