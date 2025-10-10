package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/maximberdyshev/go-study-bot/internal/roadmap"
)

type ProgressRepository struct {
	db *sql.DB
}

func NewProgressRepository(db *sql.DB) *ProgressRepository {
	return &ProgressRepository{db: db}
}

func (r *ProgressRepository) MarkDayCompleted(ctx context.Context, internalUserID int64, dayNumber int) error {
	query := `
		INSERT INTO completed_days (user_id, day_number)
		VALUES ($1, $2)
		ON CONFLICT (user_id, day_number) DO NOTHING;
	`

	if dayNumber < 1 || dayNumber > roadmap.TotalDays {
		return fmt.Errorf("day number must be between 1 and %d", roadmap.TotalDays)
	}

	_, err := r.db.ExecContext(ctx, query, internalUserID, dayNumber)
	return err
}

// future use
// func (r *UserRepository) GetCompletedDays(ctx context.Context, internalUserID int64) ([]int, error) {
// 	query := `
// 		SELECT day_number
// 		FROM completed_days
// 		WHERE user_id = $1
// 		ORDER BY day_number;
// 	`

// 	rows, err := r.db.QueryContext(ctx, query, internalUserID)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed get completed days: %w", err)
// 	}
// 	defer rows.Close()

// 	var days []int
// 	for rows.Next() {
// 		var day int
// 		if err := rows.Scan(&day); err != nil {
// 			return nil, fmt.Errorf("filed scan day: %w", err)
// 		}
// 		days = append(days, day)
// 	}

// 	return days, nil
// }

func (r *ProgressRepository) GetCompletionStats(ctx context.Context, internalUserID int64) (int, int, error) {
	query := `
		SELECT COUNT(*)
		FROM completed_days
		WHERE user_id = $1;
	`

	var completed int
	err := r.db.QueryRowContext(ctx, query, internalUserID).Scan(&completed)
	return completed, roadmap.TotalDays, err
}
