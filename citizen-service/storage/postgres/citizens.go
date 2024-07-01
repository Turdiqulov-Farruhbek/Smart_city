package postgres

import (
	citi "citizen/genproto/citizen"
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/rs/zerolog/log"
)

type Citizen struct {
	Db *pgx.Conn
}

func NewCitizen(db *pgx.Conn) *Citizen {
	return &Citizen{Db: db}
}


// RegisterCitizen inserts a new citizen into the database
func (r *Citizen) RegisterCitizen(ctx context.Context, req *citi.CitizenCreate) (*citi.Citizen, error) {
	citizenID := uuid.NewString()
	userID := uuid.NewString()

	query := `
    INSERT INTO citizens (citizen_id, user_id, first_name, last_name, date_of_birth, address, phone_number, email)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING 
        citizen_id,
        user_id,
        first_name,
        last_name,
        date_of_birth,
        address,
        phone_number,
        email,
        created_at`

	citiRes := &citi.Citizen{}
	var dateOfBirth time.Time
	var createdAt time.Time

	// Parse date_of_birth from string to time.Time
	dateOfBirth, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		slog.Error("Error parsing date_of_birth", err)
		return nil, err
	}

	err = r.Db.QueryRow(ctx, query,
		citizenID,
		userID,
		req.FirstName,
		req.LastName,
		dateOfBirth, // Pass dateOfBirth as time.Time
		req.Address,
		req.PhoneNumber,
		req.Email,
	).Scan(
		&citiRes.CitizenId,
		&citiRes.UserId,
		&citiRes.FirstName,
		&citiRes.LastName,
		&dateOfBirth, // Scan date_of_birth into time.Time
		&citiRes.Address,
		&citiRes.PhoneNumber,
		&citiRes.Email,
		&createdAt,
	)

	if err != nil {
		slog.Error("Error registering citizen", err)
		return nil, err
	}

	// Convert time.Time to string
	citiRes.DateOfBirth = dateOfBirth.Format("2006-01-02")
	citiRes.CreateAt = createdAt.Format(time.RFC3339)

	return citiRes, nil
}


// GetCitizenProfile retrieves a citizen's profile by their ID
// GetCitizenProfile retrieves a citizen's profile by their ID
func (r *Citizen) GetCitizenProfile(ctx context.Context, req *citi.ById) (*citi.Citizen, error) {
	query := `
    SELECT
        citizen_id,
        user_id,
        first_name,
        last_name,
        date_of_birth,
        address,
        phone_number,
        email,
        created_at,
        updated_at
    FROM citizens
    WHERE citizen_id = $1 AND deleted_at = 0`

	citiRes := &citi.Citizen{}
	var dateOfBirth time.Time
	var createdAt, updatedAt sql.NullTime

	err := r.Db.QueryRow(ctx, query, req.Id).Scan(
		&citiRes.CitizenId,
		&citiRes.UserId,
		&citiRes.FirstName,
		&citiRes.LastName,
		&dateOfBirth,
		&citiRes.Address,
		&citiRes.PhoneNumber,
		&citiRes.Email,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		slog.Error("Error retrieving citizen profile", err)
		return nil, err
	}

	// Convert time.Time to string
	citiRes.DateOfBirth = dateOfBirth.Format("2006-01-02")
	citiRes.CreateAt = createdAt.Time.Format(time.RFC3339)
	citiRes.UpdateAt = updatedAt.Time.Format(time.RFC3339)


	return citiRes, nil
}

func (r *Citizen) UpdateCitizenProfile(ctx context.Context, req *citi.CitizenCreate) (*citi.Citizen, error) {
	query := `
    UPDATE citizens
    SET
        first_name = $2,
        last_name = $3,
        date_of_birth = $4,
        address = $5,
        phone_number = $6,
        email = $7,
        updated_at = NOW()
    WHERE citizen_id = $1 AND deleted_at = 0
    RETURNING
        citizen_id,
        user_id,
        first_name,
        last_name,
        date_of_birth,
        address,
        phone_number,
        email,
        created_at,
        updated_at`

	citiRes := citi.Citizen{}
	var dateOfBirth time.Time
	var createdAt, updatedAt sql.NullTime
	err := r.Db.QueryRow(ctx, query,
		req.CitizenId,
		req.FirstName,
		req.LastName,
		req.DateOfBirth,
		req.Address,
		req.PhoneNumber,
		req.Email,
	).Scan(
		&citiRes.CitizenId,
		&citiRes.UserId,
		&citiRes.FirstName,
		&citiRes.LastName,
		&dateOfBirth,
		&citiRes.Address,
		&citiRes.PhoneNumber,
		&citiRes.Email,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error updating citizen profile")
		slog.Error("Error updating citizen profile", err)
		return nil, err
	}

	// Convert time.Time to string
	citiRes.DateOfBirth = dateOfBirth.Format("2006-01-02")
	citiRes.CreateAt = createdAt.Time.Format(time.RFC3339)
	citiRes.UpdateAt = updatedAt.Time.Format(time.RFC3339)
	return &citiRes, nil
}




// DeleteCitizenProfile performs a soft delete on a citizen's profile by their ID
func (r *Citizen) DeleteCitizenProfile(ctx context.Context, req *citi.ById) (*citi.Void, error) {
	query := `
    UPDATE citizens
    SET deleted_at = EXTRACT(EPOCH FROM NOW())
    WHERE citizen_id = $1 AND deleted_at = 0`

	res, err := r.Db.Exec(ctx, query, req.Id)
	if err != nil {
		slog.Error("Error performing soft delete on citizen profile", err)
		return nil, err
	}

	// Check if the update affected any rows
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		err = sql.ErrNoRows
		slog.Warn("No citizen profile found or already soft deleted", err)
		return nil, err
	}

	return &citi.Void{}, nil
}



