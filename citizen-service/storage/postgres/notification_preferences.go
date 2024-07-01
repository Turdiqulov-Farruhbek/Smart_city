package postgres

import (
	"context"
	"database/sql"
	"log/slog"

	citi "citizen/genproto/citizen"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/rs/zerolog/log"
)

// Notification handles CRUD operations for Notification Preferences
type Notification struct {
	db *pgx.Conn
}

// NewNotificationRepository initializes a new Notification
func NewNotification(db *pgx.Conn) *Notification {
	return &Notification{db: db}
}

// SetNotificationPref inserts a new notification preference into the database
func (r *Notification) SetNotificationPref(ctx context.Context, req *citi.NotificationPref) (*citi.Void, error) {
	PreferenceId := uuid.NewString()
	query := `
    INSERT INTO notification_preferences (preference_id, citizen_id, notification_type, is_enabled)
    VALUES ($1, $2, $3, $4)`

	_, err := r.db.Exec(ctx, query,
		PreferenceId,
		req.CitizenId,
		req.NotificationType,
		req.IsEnabled,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error setting notification preference")
		slog.Error("Error setting notification preference", err)
		return nil, err
	}

	return &citi.Void{}, nil
}

// UpdateNotificationPref updates an existing notification preference in the database
func (r *Notification) UpdateNotificationPref(ctx context.Context, req *citi.NotificationPref) (*citi.Void, error) {
	query := `
    UPDATE notification_preferences
    SET
        notification_type = $2,
        is_enabled = $3,
        update_at = NOW()
    WHERE preference_id = $1 AND deleted_at = 0`

	_, err := r.db.Exec(ctx, query,
		req.PreferenceId,
		req.NotificationType,
		req.IsEnabled,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error updating notification preference")
		slog.Error("Error updating notification preference", err)
		return nil, err
	}

	return &citi.Void{}, nil
}

// DeleteNotificationPref deletes a notification preference from the database
func (r *Notification) DeleteNotificationPref(ctx context.Context, req *citi.ById) (*citi.Void, error) {
	query := `
    UPDATE notification_preferences
    SET delete_at = EXTRACT(EPOOCH FROM NOW()
    WHERE preference_id = $1 AND delete_at = 0`

	res, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Error performing soft delete on notification preference")
		slog.Error("Error performing soft delete on notification preference", err)
		return nil, err
	}

	// Check if the update affected any rows
	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		err = sql.ErrNoRows
		log.Warn().Err(err).Msg("No notification preference found or already soft deleted")
		slog.Warn("No notification preference found or already soft deleted", err)
		return nil, err
	}

	return &citi.Void{}, nil
}
