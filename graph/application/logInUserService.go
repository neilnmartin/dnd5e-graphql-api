package application

import (
	"log"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"golang.org/x/crypto/bcrypt"
)

// LogInResponse consists of the logged in User as well as the token
type LogInResponse struct {
	User  *domain.User
	token string
}

// LogInUserService logs in a user by comparing password hashes and returning a token
func LogInUserService(email string, password string) (*LogInResponse, error) {
	log.Println("reached login service")

	dbu, err := db.UserRepo.GetUserByEmail(email)
	if err != nil {
		log.Printf("%v", err)
		return nil, err
	}

	hpw := []byte(dbu.Password)
	err = bcrypt.CompareHashAndPassword(hpw, []byte(password))
	if err != nil {
		log.Printf("%v", err) // throw error
		return nil, err
	}

	return &LogInResponse{
		User:  dbu,
		token: "token",
	}, nil
}
