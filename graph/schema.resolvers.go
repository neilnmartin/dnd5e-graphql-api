package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/application"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/generated"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/repository"
)

func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.CreateCharacterInput) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignUpUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	ur := repository.DB.UserRepo
	dn := domain.Name{
		GivenName:  input.Name.GivenName,
		FamilyName: input.Name.FamilyName,
	}
	u, err := application.SignUpUserService(input.Email, input.Password, dn, ur)
	if err != nil {
		return nil, err
	}
	// convert from domain to api model
	mu := &model.User{
		Email: u.Email,
		Name: &model.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
	}
	return mu, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	lu, err := application.LogInUserService(input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	return &model.LoginResponse{
		User: &model.User{
			Email: lu.User.Email,
			Name: &model.Name{
				GivenName:  lu.User.GivenName,
				FamilyName: lu.User.FamilyName,
			},
		},
		// Token: l.Token,
	}, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Races(ctx context.Context) ([]*model.Race, error) {
	ar, err := repository.DB.RaceRepo.GetAllRaces()
	if err != nil {
		return nil, err
	}
	rm := []*model.Race{}
	for _, dr := range *ar {
		fmt.Printf("\n%v %v", dr.Name, dr.Size)
		rm = append(rm, &model.Race{
			ID:                  dr.ID,
			Name:                dr.Name,
			Age:                 &dr.Age,
			Size:                &dr.Size,
			SizeDescription:     &dr.SizeDescription,
			Speed:               &dr.Speed,
			LanguageDescription: &dr.LanguageDescription,
		})
	}
	return rm, nil
}

func (r *queryResolver) Classes(ctx context.Context) ([]*model.Class, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
