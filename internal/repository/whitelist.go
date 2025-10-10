package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type WhitelistRepository struct {
	db *sql.DB
}

func NewWhitelistRepository(db *sql.DB) *WhitelistRepository {
	return &WhitelistRepository{db: db}
}

func (r *WhitelistRepository) IsAllowed(ctx context.Context, telegramID int64) (bool, error) {
	query := `
		SELECT true
		FROM allowed_users
		WHERE telegram_id = $1;
	`

	var exists bool
	if err := r.db.QueryRowContext(ctx, query, telegramID).Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed check allowed_users: %w", err)
	}
	return true, nil
}
