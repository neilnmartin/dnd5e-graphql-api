package repository

import (
	"gopkg.in/mgo.v2"
)

// A Datasource that implements a data repository with CRUD for each resource
type Datasource struct {
	UserRepo userMongoRepo
	RaceRepo raceMongoRepo
}

// DatasourceFactory creates a Datasource based on a string input p
// p will identify the persistence infrastructure used
// eg. in this case "mongodb"
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
	umr := userMongoRepo{
		session: session,
	}
	rmr := raceMongoRepo{
		session: session,
	}
	return Datasource{
		UserRepo: umr,
		RaceRepo: rmr,
	}
}

// DB is the data source instance
var DB = DatasourceFactory("mongodb")
