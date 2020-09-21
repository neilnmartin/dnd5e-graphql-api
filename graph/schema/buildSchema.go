package schema

import (
	"fmt"
)

/*
since the schema.graphqls file is required to generate with gqlgen,
I am implementing a schema builder that will build a schema.graphqls file
from multiple schema files
*/
func buildSchema() int {
	fmt.Println("will build schema here")
	// read and compile into a single schema.graphqls file
	return 0
}

var success = buildSchema()
