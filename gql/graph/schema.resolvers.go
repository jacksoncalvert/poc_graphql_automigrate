package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"log"

	"context"
	"gql/graph/generated"
	"gql/graph/model"
	gApi "gql/api"
)

// GetHellOr is the resolver for the getHellOr field.
func (r *queryResolver) GetHellOr(ctx context.Context) (*model.HellOr, error) {
	resp, resp2, err := gApi.QueryService()
	log.Println(resp2)
	if err != nil {
		log.Fatalln(err)
	}
	result := model.HellOr {
		ID: resp,
		Msg: resp2,
	}

	return &result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
