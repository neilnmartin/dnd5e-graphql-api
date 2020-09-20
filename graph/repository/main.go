package repository

// A datasource that implements a data repository with CRUD for each resource
type datasource struct {
	UserRepo userMongoRepo
	RaceRepo raceMongoRepo
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
var DB = datasourceFactory("mongodb")
