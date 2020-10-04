package application

import (
	"errors"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
)

//CreateCharacterInput is the input for the CreateCharacterService
type CreateCharacterInput struct {
	Race       string
	Class      string
	SubClass   string
	SubRace    string
	Background string
	UserID     string
	Name       string
}

// CreateCharacterService signs up a user by adding to the database after validating name email and password inputs
func CreateCharacterService(input CreateCharacterInput, cr domain.CharacterRepo) (*domain.Character, error) {
	if &input.UserID != nil {
		return nil, errors.New("Invalid user id")
	}
	u := domain.User{ID: input.UserID}
	c := domain.Character{
		Name:     input.Name,
		Race:     domain.Race{Name: input.Race},
		SubRace:  domain.SubRace{Name: input.SubRace},
		Class:    domain.Class{Name: input.Class},
		SubClass: domain.SubClass{Name: input.SubClass},
	}
	created, err := cr.InsertCharacter(c, u)
	if err != nil {
		if err == domain.ErrCharacterNotFound {
			return nil, err
		}
	}
	return created, nil
}
