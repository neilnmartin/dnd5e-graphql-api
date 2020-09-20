package application

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"golang.org/x/crypto/bcrypt"
)

// SignUpUserService signs up a user by adding to the database after validating name email and password inputs
func SignUpUserService(email string, password string, name domain.Name, ur domain.UserRepo) (*domain.User, error) {
	// check existing user: TODO
	eu, err := ur.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if eu != nil {
		return nil, errors.New("A user with this email already exists")
	}

	validEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !validEmail.MatchString(email) {
		log.Fatal("not a valid email")
		return nil, errors.New("Invalid email")
	}

	hpw, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	if err != nil {
		log.Printf(err.Error())
	}

	shpw := string(hpw)
	fmt.Printf("%v original pw, %v stringified hashed pw", password, shpw)

	nu := domain.User{
		Email:    email,
		Password: shpw,
		Name: domain.Name{
			GivenName:  name.GivenName,
			FamilyName: name.FamilyName,
		},
	}

	user, err := ur.CreateUser(nu)
	return user, nil
}
