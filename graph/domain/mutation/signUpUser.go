package mutation

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// Name consists of the givenName, familyName
type Name struct {
	givenName  string
	familyName string
}

// SignUpUser signs up a user by adding to the database after validating name email and password inputs
func SignUpUser(email string, password string, name Name) {
	// TODO
	// add userRepo helper
	hpw, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		log.Printf(err.Error())
	}
	fmt.Printf("%v", hpw)
}
