package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/application"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/domain"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/generated"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/repository"
	"github.com/neilnmartin/dnd5e-graphql-api/utils"
)

func (r *mutationResolver) CreateCharacter(ctx context.Context, input model.CreateCharacterInput) (*model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SignUpUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	log.Printf("hit signup mutation resolver")
	ur := repository.DB.UserRepo
	dn := domain.Name{
		GivenName:  input.Name.GivenName,
		FamilyName: input.Name.FamilyName,
	}
	u, err := application.SignUpUserService(input.Email, input.Password, dn, ur)
	if err != nil {
		return nil, err
	}
	log.Printf("\nResolver before model convert: %+v", u)
	// convert from domain to api model
	mu := &model.User{
		ID:    u.ID,
		Email: u.Email,
		Name: &model.Name{
			GivenName:  u.GivenName,
			FamilyName: u.FamilyName,
		},
	}
	log.Printf("\nResolver after model convert: %+v", mu)
	return mu, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginInput) (*model.LoginResponse, error) {
	lu, err := application.LogInUserService(input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	idToken, err := utils.GenerateJWT(utils.IDToken{
		ID: lu.ID,
		Name: model.Name{
			GivenName:  lu.Name.GivenName,
			FamilyName: lu.Name.FamilyName,
		},
	})
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		User: &model.User{
			ID:    lu.ID,
			Email: lu.Email,
			Name: &model.Name{
				GivenName:  lu.GivenName,
				FamilyName: lu.FamilyName,
			},
		},
		Token: *idToken,
	}, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Class(ctx context.Context, name string) (*model.Class, error) {
	cd, err := repository.DB.ClassRepo.GetClassByName(name)
	if err != nil {
		return nil, err
	}
	return mapClassFromDomainToAPI(*cd), nil
}

func (r *queryResolver) Race(ctx context.Context, name string) (*model.Race, error) {
	cd, err := repository.DB.RaceRepo.GetRaceByName(name)
	if err != nil {
		return nil, err
	}
	return mapRaceFromDomainToAPI(*cd), nil
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
			ID:                  &dr.ID,
			Name:                &dr.Name,
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
	log.Println("hit classes query resolver")
	ac, err := repository.DB.ClassRepo.GetAllClasses()
	if err != nil {
		return nil, err
	}
	mc := []*model.Class{}
	for _, dc := range *ac {
		log.Printf("%v", dc)
		fmt.Printf("\n%v", dc.Name)
		mc = append(mc, &model.Class{
			ID:     &dc.ID,
			Name:   &dc.Name,
			HitDie: &dc.HitDie,
		})
	}
	return mc, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
