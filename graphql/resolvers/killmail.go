package resolvers

import (
	"context"

	"github.com/ddouglas/neo"
	"github.com/ddouglas/neo/graphql/service"
)

func (r *queryResolver) Killmail(ctx context.Context, id int) (*killboard.Killmail, error) {
	return r.KillmailServ.Killmail(ctx, uint64(id))
}

func (r *Resolver) Killmail() service.KillmailResolver {
	return &killmailResolver{r}
}

type killmailResolver struct{ *Resolver }

func (r *killmailResolver) Attackers(ctx context.Context, obj *killboard.Killmail) ([]*killboard.KillmailAttacker, error) {
	panic("not implemented")
}
func (r *killmailResolver) Victim(ctx context.Context, obj *killboard.Killmail) (*killboard.KillmailVictim, error) {
	panic("not implemented")
}
