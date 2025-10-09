package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(ctx context.Context, dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
