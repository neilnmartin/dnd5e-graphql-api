package repository

import (
	"crypto/tls"
	"net"

	"github.com/neilnmartin/dnd5e-graphql-api/config"
	"gopkg.in/mgo.v2"
)

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
	// mongo atlas rejects unsecured connections
	di.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), &tls.Config{})
		return conn, err
	}
	return mgo.DialWithInfo(&di)
}

func createMongoDataSource() datasource {
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
	cmr := classMongoRepo{
		session: session,
	}
	chmr := characterMongoRepo{
		session: session,
	}
	return datasource{
		UserRepo:      umr,
		RaceRepo:      rmr,
		ClassRepo:     cmr,
		CharacterRepo: chmr,
	}
}
