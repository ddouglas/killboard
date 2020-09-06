package mdb

import (
	"context"
	"time"

	"github.com/eveisesi/neo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type allianceRepository struct {
	c *mongo.Collection
}

func NewAllianceRepository(d *mongo.Database) neo.AllianceRespository {
	return &allianceRepository{
		d.Collection("alliances"),
	}
}

func (r *allianceRepository) Alliance(ctx context.Context, id uint) (*neo.Alliance, error) {

	var alliance = new(neo.Alliance)

	err := r.c.FindOne(ctx, primitive.D{primitive.E{Key: "id", Value: id}}).Decode(alliance)
	return alliance, err

}
func (r *allianceRepository) Alliances(ctx context.Context, mods ...neo.Modifier) ([]*neo.Alliance, error) {

	pds := BuildFilters(mods...)
	pos := BuildFindOptions(mods...)

	var alliances = make([]*neo.Alliance, 0)
	result, err := r.c.Find(ctx, pds, pos)
	if err != nil {
		return nil, err
	}

	err = result.All(ctx, &alliances)
	return alliances, err

}

func (r *allianceRepository) CreateAlliance(ctx context.Context, alliance *neo.Alliance) error {

	alliance.CreatedAt = time.Now().Unix()
	alliance.UpdatedAt = time.Now().Unix()

	_, err := r.c.InsertOne(ctx, alliance)
	if err != nil {
		if !IsUniqueConstrainViolation(err) {
			return err
		}
	}

	return nil

}
func (r *allianceRepository) UpdateAlliance(ctx context.Context, id uint, alliance *neo.Alliance) error {

	alliance.UpdatedAt = time.Now().Unix()
	if alliance.CreatedAt == 0 {
		alliance.CreatedAt = time.Now().Unix()
	}

	update := primitive.D{primitive.E{Key: "$set", Value: alliance}}

	_, err := r.c.UpdateOne(ctx, primitive.D{{Key: "id", Value: id}}, update, nil)

	return err

}

func (r *allianceRepository) Expired(ctx context.Context) ([]*neo.Alliance, error) {

	mods := []neo.Modifier{
		neo.LessThan{Column: "cachedUntil", Value: time.Now().Unix()},
		neo.LimitModifier(1000),
		neo.OrderModifier{Column: "cachedUntil", Sort: neo.SortAsc},
	}

	return r.Alliances(ctx, mods...)

}
