package repository

import (
	"context"
	"fmt"
)

func (r *UserRepository) InitSchema(ctx context.Context) error {
	schema := `
  CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    username TEXT,
    first_name TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );
  `

	_, err := r.db.ExecContext(ctx, schema)
	if err != nil {
		return fmt.Errorf("failed to init schema: %w", err)
	}
	return nil
}
