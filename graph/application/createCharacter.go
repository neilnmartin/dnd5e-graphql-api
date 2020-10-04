package application

import (
	"errors"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
)

//CreateCharacterInput is the input for the CreateCharacterService
type CreateCharacterInput struct {
	race       string
	class      string
	subClass   string
	subRace    string
	background string
	uid        string
	name       string
	cr         domain.CharacterRepo
}

// CreateCharacterService signs up a user by adding to the database after validating name email and password inputs
func CreateCharacterService(input CreateCharacterInput) (*domain.Character, error) {
	if &input.uid != nil {
		return nil, errors.New("Invalid user id")
	}
	u := domain.User{ID: input.uid}
	c := domain.Character{
		Name:     input.name,
		Race:     domain.Race{Name: input.race},
		SubRace:  domain.SubRace{Name: input.subRace},
		Class:    domain.Class{Name: input.class},
		SubClass: domain.SubClass{Name: input.subClass},
	}
	created, err := input.cr.InsertCharacter(c, u)
	if err != nil {
		if err == domain.ErrCharacterNotFound {
			return nil, err
		}
	}
	return created, nil
}
