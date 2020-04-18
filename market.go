package neo

import (
	"context"
	"time"
)

type MarketRepository interface {
	Orders(context.Context, uint64) ([]*Order, error)
	CreateOrdersBulk(ctx context.Context, orders []*Order) ([]*Order, error)
}

type Order struct {
	TypeID    uint      `json:"typeID"`
	Date      time.Time `json:"date"`
	LowPrice  float64   `json:"lowPrice"`
	HighPrice float64   `json:"highPrice"`
	AvgPrice  float64   `json:"avgPrice"`
	Tenfold   bool      `json:"tenfold"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
