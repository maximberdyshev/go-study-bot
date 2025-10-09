package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type User struct {
	ID         int64  `json:"id"`
	TelegramID int64  `json:"telegram_id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUserIfNotExists(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (telegram_id, username, first_name)
	VALUES ($1, $2, $3)
	ON CONFLICT (telegram_id) DO NOTHING;
	`

	_, err := r.db.ExecContext(ctx, query, user.TelegramID, user.Username, user.FirstName)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}
