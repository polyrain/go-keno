package db

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/polyrain/go-keno/pkg/keno"
)

type PgPool struct {
	Pool *pgxpool.Pool
}

// Bail out on failure to connext for now
func NewPgPool() *PgPool {

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("failed to connect to postgres, check your url", "err", err)
		os.Exit(1)
	}

	return &PgPool{Pool: dbpool}
}

func (pool *PgPool) GetGame(gameID int) (keno.Game, error) {
	// Insert DB query here
	return keno.Game{}, nil
}

func (pool *PgPool) GetCard(cardId int) (keno.Card, error) {
	// Insert ANOTHER DB query here
	return keno.Card{}, nil
}

func (pool *PgPool) GetGamesBetween(startId, endId int) ([]keno.Game, error) {
	// Insert the MOST COMPLEX query here
	return nil, nil
}
