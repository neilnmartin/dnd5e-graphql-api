package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/neilnmartin/dnd5e-graphql-api/graph/generated"
	"github.com/neilnmartin/dnd5e-graphql-api/graph/model"
)

func (r *classResolver) SubClasses(ctx context.Context, obj *model.Class) ([]*model.SubClass, error) {
	panic(fmt.Errorf("not implemented"))
}

// Class returns generated.ClassResolver implementation.
func (r *Resolver) Class() generated.ClassResolver { return &classResolver{r} }

type classResolver struct{ *Resolver }
