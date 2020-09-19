package application

import (
	"fmt"
	"log"
	"regexp"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"golang.org/x/crypto/bcrypt"
)

// Name consists of the givenName, familyName
type Name struct {
	givenName  string
	familyName string
}

// SignUpUser signs up a user by adding to the database after validating name email and password inputs
func SignUpUser(email string, password string, name Name, userRepo domain.UserRepo) domain.User {
	// check existing user: TODO
	validEail := regexp.MustCompile(``)

	if !validEail.MatchString(email) {
		log.Fatal("not a valid email")
	}

	hpw, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		log.Printf(err.Error())
	}

	shpw := string(hpw)
	fmt.Printf("%v original pw, %v stringified hashed pw", password, shpw)

	nu := domain.User{
		Email:    email,
		Password: string(hpw),
		Name: domain.Name{
			GivenName:  name.givenName,
			FamilyName: name.familyName,
		},
	}

	user := db.UserRepo.CreateUser(nu)
	return user
}
