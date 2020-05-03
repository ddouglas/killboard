package resolvers

import (
	"context"

	"github.com/eveisesi/neo"
	"github.com/eveisesi/neo/graphql/service"
)

func (r *queryResolver) CorporationByCorporationID(ctx context.Context, id int) (*neo.Corporation, error) {
	return r.Services.Corporation.Corporation(ctx, uint64(id))
}

func (r *queryResolver) KillmailsByCorporationID(ctx context.Context, corporationID int, page *int) ([]*neo.Killmail, error) {
	return r.Services.KillmailsByCorporationID(ctx, uint64(corporationID), *page)
}

func (r *Resolver) Corporation() service.CorporationResolver {
	return &corporationResolver{r}
}

type corporationResolver struct{ *Resolver }

func (r *corporationResolver) Alliance(ctx context.Context, obj *neo.Corporation) (*neo.Alliance, error) {
	if !obj.AllianceID.Valid {
		return nil, nil
	}
	return r.Dataloader(ctx).AllianceLoader.Load(obj.AllianceID.Uint64)
}
