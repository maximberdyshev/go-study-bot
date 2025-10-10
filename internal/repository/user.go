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
		ON CONFLICT (telegram_id) DO UPDATE SET
			username = EXCLUDED.username,
			first_name = EXCLUDED.first_name;
	`

	_, err := r.db.ExecContext(ctx, query, user.TelegramID, user.Username, user.FirstName)
	return err
}

func (r *UserRepository) GetUserByTelegramID(ctx context.Context, telegramID int64) (*User, error) {
	query := `
		SELECT id, telegram_id, username, first_name
		FROM users
		WHERE telegram_id = $1;
	`

	var user User
	if err := r.db.QueryRowContext(ctx, query, telegramID).Scan(
		&user.ID,
		&user.TelegramID,
		&user.Username,
		&user.FirstName,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with telegram_id %d not found", telegramID)
		}
		return nil, fmt.Errorf("failed get user: %w", err)
	}
	return &user, nil

}
