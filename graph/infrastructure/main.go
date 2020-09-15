package infrastructure

import (
	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"gopkg.in/mgo.v2"
)

// A Datasource that implements a data repository with CRUD for each resource
type Datasource interface {
	CreateUser(domain.User) domain.User
	// ReadUser(domain.User) domain.User
}

// MongoUserRepo is the mongodb infrastructure implementation of the UserRepo domain interface
type MongoUserRepo interface {
	CreateUser(domain.User) domain.User
	GetUserByEmail(domain.User) domain.User
	UpdateUser(domain.User) domain.User
	DeleteUser(domain.User) domain.User
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

func createMongoDataSource() mongoDatasource {

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	return mongoDatasource{
		session,
	}
}
