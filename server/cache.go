package server

import (
	"context"
	"fmt"
	"time"

	"github.com/eveisesi/neo"
	"github.com/go-redis/redis/v7"
)

type GQLCache struct {
	client *redis.Client
	ttl    time.Duration
}

func (c *GQLCache) Add(ctx context.Context, hash string, query string) {
	c.client.Set(fmt.Sprintf("%s:%s", neo.REDIS_GRAPHQL_APQ_CACHE, hash), query, c.ttl)
}

func (c *GQLCache) Get(ctx context.Context, hash string) (string, bool) {

	s, err := c.client.Get(hash).Result()
	if err != nil {
		return "", false
	}

	return s, true
}
