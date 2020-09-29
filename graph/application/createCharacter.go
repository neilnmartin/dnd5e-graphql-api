package application

import (
	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
)

//CreateCharacterInput is the input for the CreateCharacterService
type CreateCharacterInput struct {
	race       string
	class      string
	subclass   string
	background string
	name       domain.Name
	ur         domain.UserRepo
	cr         domain.CharacterRepo
}

// CreateCharacterService signs up a user by adding to the database after validating name email and password inputs
func CreateCharacterService() (*domain.Character, error) {
	return &domain.Character{
		Name: "",
	}, nil
}
