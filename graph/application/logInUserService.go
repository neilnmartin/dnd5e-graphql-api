package application

import (
	"errors"
	"log"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"golang.org/x/crypto/bcrypt"
)

// LogInResponse consists of the logged in User as well as the token
type LogInResponse struct {
	User  *domain.User
	Token string
}

// LogInUserService logs in a user by comparing password hashes and returning a token
func LogInUserService(email string, password string) (*domain.User, error) {
	log.Println("reached login service")

	dbu, err := db.UserRepo.GetUserByEmail(email)
	if err != nil {
		if err.Error() == "not found" {
			return nil, errors.New("User does not exist")
		}
		log.Printf("\nLogin Error: %+v", err)
		return nil, err
	}

	hpw := []byte(dbu.Password)
	err = bcrypt.CompareHashAndPassword(hpw, []byte(password))
	if err != nil {
		log.Printf("\nHash compare error: %+v", err) // throw error
		return nil, errors.New("Incorrect password")
	}

	return dbu, nil
}
