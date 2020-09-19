package graph

import (
	"context"
	"fmt"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
}

// Character finds a character
func (r *Resolver) Character(ctx context.Context) (*model.Character, error) {
	fmt.Printf("%v", "character resolver hit")
	panic(fmt.Errorf("not implemented"))
}

// func (r *Resolver) Race(ctx context.Context) (*model.Race, error) {
// 	var rr domain.Race
// 	if input.id != nil {
// 		rr = repository.DB.GetRaceById(input.id)
// 		return race, nil
// 	} else if input.name != nil {
// 		rr = repository.DB.GetRaceByName(input.name)
// 		return race, nil
// 	} else {
// 		return nil, errors.New("Could not find Race")
// 	}
// }
// func (r *Resolver) Class(ctx context.Context) (*model.Class, error) {
// 	panic(fmt.Errorf("not implemented"))
// }
// func (r *Resolver) Spell(ctx context.Context) (*model.Spell, error) {
// 	panic(fmt.Errorf("not implemented"))
// }
