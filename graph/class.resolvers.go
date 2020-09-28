package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/generated"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/repository"
)

func (r *classResolver) SubClasses(ctx context.Context, obj *model.Class) ([]*model.SubClass, error) {
	scm := []*model.SubClass{}
	for _, osc := range obj.SubClasses {
		sc, err := repository.DB.ClassRepo.GetSubClassByName(*osc.Name)
		if err != nil {
			return nil, err
		}
		scm = append(scm, &model.SubClass{
			ID:          &sc.ID,
			Name:        &sc.Name,
			Description: &sc.Description,
			Flavor:      &sc.Flavor,
		})
	}

	return scm, nil
}

// Class returns generated.ClassResolver implementation.
func (r *Resolver) Class() generated.ClassResolver { return &classResolver{r} }

type classResolver struct{ *Resolver }
