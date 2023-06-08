package repository

import (
	"context"
	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"net/url"
)

func NewPostgresDB(ctx context.Context) (*pgxpool.Pool, *pgadapter.Adapter, error) {
	dbURL := "postgres://" + "postgres" + ":" + url.QueryEscape("password0701") + "@" + "casbin-db" + ":" + "5432" + "/" + "practice" + "?sslmode=" + "disable"
	dbPool, err := pgxpool.Connect(ctx, dbURL)

	if err != nil {
		return nil, nil, err
	}
	adapter, err := pgadapter.NewAdapter(dbURL, "practice")
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to open adapter connection to database")
	}
	err = dbPool.Ping(ctx)
	if err != nil {
		return nil, nil, err
	}
	return dbPool, adapter, nil
}
