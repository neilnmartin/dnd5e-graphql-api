package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/application"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/repository"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/generated"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
)

func (r *mutationResolver) SignUpUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	sampleUser := model.User{
		ID:    "sampleid",
		Email: "sampleemail@email.com",
	}
	return &sampleUser, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	l := application.LogInUser(input.Email, input.Password)
	return &model.LoginResponse{
		User: &model.User{
			Email: l.User.Email,
		},
		// Token: l.Token,
	}, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Character(ctx context.Context) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Race(ctx context.Context) (*model.Race, error) {
	var race domain.Race
	if input.id != nil {
		race = repository.DB.GetRaceById(input.id)
		return race, nil
	} else if input.name != nil {
		race = repository.DB.GetRaceByName(input.name)
		return race, nil
	} else {
		return nil, errors.New("Could not find Race")
	}
}

func (r *queryResolver) Class(ctx context.Context) (*model.Class, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Spell(ctx context.Context) (*model.Spell, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
