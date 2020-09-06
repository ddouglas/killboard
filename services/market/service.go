package market

import (
	"context"
	"time"

	"github.com/eveisesi/neo"
	"github.com/eveisesi/neo/services/esi"
	"github.com/eveisesi/neo/services/tracker"
	"github.com/eveisesi/neo/services/universe"
	"github.com/go-redis/redis/v7"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type Service interface {
	FetchHistory(ctx context.Context)
	FetchTypePrice(id uint, date time.Time) float64
	FetchPrices(ctx context.Context)
	neo.MarketRepository
}

type service struct {
	redis    *redis.Client
	esi      esi.Service
	nr       *newrelic.Application
	logger   *logrus.Logger
	universe universe.Service
	txn      neo.Starter
	neo.MarketRepository
	tracker tracker.Service
}

func NewService(redis *redis.Client, esi esi.Service, nr *newrelic.Application, logger *logrus.Logger, universe universe.Service, txn neo.Starter, market neo.MarketRepository, tracker tracker.Service) Service {
	return &service{
		redis,
		esi,
		nr,
		logger,
		universe,
		txn,
		market,
		tracker,
	}
}
