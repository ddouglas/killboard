package resolvers

import (
	"context"

	"github.com/ddouglas/killboard/graphql/service"
	"github.com/ddouglas/killboard/services/killmail"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	KillmailServ killmail.Service
}

func (r *Resolver) Mutation() service.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() service.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) MutationPlaceholder(ctx context.Context) (bool, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) QueryPlaceholder(ctx context.Context) (bool, error) {
	panic("not implemented")
}
