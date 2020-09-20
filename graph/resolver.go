package graph

import (
	"context"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var db = repository.DB

// Resolver struct holds general resolvers
type Resolver struct {
}

// Character finds a character
// func (r *Resolver) Character(ctx context.Context) (*model.Character, error) {
// 	fmt.Printf("%v", "character resolver hit")
// 	panic(fmt.Errorf("not implemented"))
// }

//RaceInput for Race resolver
type RaceInput struct {
	name string `json:"name"`
}

//Race generic resolver
func (r *Resolver) Race(ctx context.Context, input RaceInput) (*model.Race, error) {
	race, err := repository.DB.RaceRepo.GetRaceByName(input.name)
	if err != nil {
		return nil, err
	}
	return &model.Race{
		ID:                  race.ID,
		Name:                race.Name,
		Age:                 &race.Age,
		Size:                &race.Size,
		SizeDescription:     &race.SizeDescription,
		Speed:               &race.Speed,
		LanguageDescription: &race.LanguageDescription,
		// StartingProficiencies
		// Traits
		// SubRaces
	}, nil
}

// func (r *Resolver) Class(ctx context.Context) (*model.Class, error) {
// 	panic(fmt.Errorf("not implemented"))
// }
// func (r *Resolver) Spell(ctx context.Context) (*model.Spell, error) {
// 	panic(fmt.Errorf("not implemented"))
// }
