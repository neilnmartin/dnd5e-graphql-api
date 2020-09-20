package repository

import (
	"crypto/tls"
	"net"

	"github.com/neilnmartin/dnd5e-graphql-api/config"
	"gopkg.in/mgo.v2"
)

// A Datasource that implements a data repository with CRUD for each resource
type Datasource struct {
	UserRepo userMongoRepo
	RaceRepo raceMongoRepo
}

// DatasourceFactory creates a Datasource based on a string input p.
// p will identify the db or other persistence infrastructure used.
func DatasourceFactory(p string) Datasource {
	// only one implementation so far, default to mongo
	switch p {
	case "mongodb":
		return createMongoDataSource()
	default:
		return createMongoDataSource()
	}
}

func createMongoDataSource() Datasource {
	session, err := createMongoConnection()
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

func createMongoConnection() (*mgo.Session, error) {
	di := mgo.DialInfo{
		Addrs: []string{
			config.Config.MongoConfig.MongoClusterAddr1,
			config.Config.MongoConfig.MongoClusterAddr2,
			config.Config.MongoConfig.MongoClusterAddr3,
		},
		Username: config.Config.MongoConfig.MongoUser,
		Password: config.Config.MongoConfig.MongoPassword,
	}
	// atlas rejects unsecured connections
	di.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), &tls.Config{})
		return conn, err
	}
	return mgo.DialWithInfo(&di)
}

// DB is the data source instance
var DB = DatasourceFactory("mongodb")
