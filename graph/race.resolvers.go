package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/generated"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/repository"
)

func (r *raceResolver) SubRaces(ctx context.Context, obj *model.Race) ([]*model.SubRace, error) {
	scm := []*model.SubRace{}
	for _, osc := range obj.SubRaces {
		sc, err := repository.DB.RaceRepo.GetSubRaceByName(osc.Name)
		if err != nil {
			return nil, err
		}
		scm = append(scm, &model.SubRace{
			ID:   sc.ID,
			Race: obj,
			Name: sc.Name,
		})
	}

	return scm, nil
}

func (r *raceResolver) Traits(ctx context.Context, obj *model.Race) ([]*model.Trait, error) {
	mts := []*model.Trait{}
	for _, tm := range obj.Traits {
		t, err := repository.DB.RaceRepo.GetTraitByName(*tm.Name)
		if err != nil {
			return nil, err
		}
		mapped := mapTraitFromDomainToAPI(*t)
		mts = append(mts, mapped)
	}
	return mts, nil
}

// Race returns generated.RaceResolver implementation.
func (r *Resolver) Race() generated.RaceResolver { return &raceResolver{r} }

type raceResolver struct{ *Resolver }
