package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/YurcheuskiRadzivon/telemetrics-back/internal/infrastructure/database/queries"
)

type Database struct {
	Pool    *pgxpool.Pool
	Queries *queries.Queries
}

func New(ctx context.Context, connString string) (*Database, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, err
	}

	return &Database{
		Pool:    pool,
		Queries: queries.New(pool),
	}, nil
}

func (db *Database) Close() {
	db.Pool.Close()
}
