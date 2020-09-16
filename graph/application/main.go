package application

import (
	"github.com/neilnmartin/dnd5e-graphql-api/graph/infrastructure"
)

var repos = infrastructure.DatasourceFactory("mongodb")
var userRepo = repos.MongoUserRepo

func main() {

}
