package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/generated"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
)

func (r *raceResolver) SubRaces(ctx context.Context, obj *model.Race) ([]*model.SubRace, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *raceResolver) Traits(ctx context.Context, obj *model.Race) ([]*model.Trait, error) {
	panic(fmt.Errorf("not implemented"))
}

// Race returns generated.RaceResolver implementation.
func (r *Resolver) Race() generated.RaceResolver { return &raceResolver{r} }

type raceResolver struct{ *Resolver }
