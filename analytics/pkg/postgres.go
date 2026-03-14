package pkg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupPostgres(ctx context.Context, url string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, err
	}

	return conn, nil
}