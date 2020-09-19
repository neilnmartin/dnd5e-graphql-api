package application

import (
	"fmt"
	"log"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"golang.org/x/crypto/bcrypt"
)

// LogInResponse consists of the logged in User as well as the token
type LogInResponse struct {
	User  *domain.User
	token string
}

// LogInUser logs in a user by comparing password hashes and returning a token
func LogInUser(email string, password string) LogInResponse {
	fmt.Println("reached login service")

	dbu := db.UserRepo.GetUserByEmail(email)
	hpw := []byte(dbu.Password)
	err := bcrypt.CompareHashAndPassword(hpw, []byte(password))
	if err != nil {
		log.Printf("%v", err) // throw error
	}

	return LogInResponse{
		User:  &dbu,
		token: "token",
	}
}
