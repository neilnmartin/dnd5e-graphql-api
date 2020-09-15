package infrastructure

import (
	"gopkg.in/mgo.v2"
)

// A Datasource that implements a data repository with CRUD for each resource
type Datasource interface {
	MongoUserRepo
}
type mongoDatasource struct {
	session *mgo.Session
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
