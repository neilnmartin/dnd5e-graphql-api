package repository

import (
	"gopkg.in/mgo.v2"
)

// A Datasource that implements a data repository with CRUD for each resource
type Datasource struct {
	UserRepo mongoUserRepo
	RaceRepo raceMongoRepo
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

func createMongoDataSource() Datasource {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	mur := mongoUserRepo{
		session: session,
	}
	return Datasource{
		UserRepo: mur,
	}
}

// DB is the data source instance
var DB = DatasourceFactory("mongodb")
