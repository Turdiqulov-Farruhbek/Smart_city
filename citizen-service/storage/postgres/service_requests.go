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

// ServiceRequest handles CRUD operations for Citizen Service Requests
type ServiceRequest struct {
	db *pgx.Conn
}

// NewCitizenRequestRepository initializes a new ServiceRequest
func NewServiceRequest(db *pgx.Conn) *ServiceRequest {
	return &ServiceRequest{db: db}
}

// CreateServiceRequest creates a new service request in the database
func (r *ServiceRequest) CreateServiceRequest(ctx context.Context, req *citi.ServiceReqCreate) (*citi.ServiceReq, error) {
	RequestId := uuid.NewString()
	query := `
    INSERT INTO service_requests (request_id, citizen_id, request_type, description, status)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING 
        RequestId,
        CitizenId,
        RequestType,
        Description,
        Status,
        CreateAt`

	serviceReq := citi.ServiceReq{}
	err := r.db.QueryRow(ctx, query,
		RequestId,
		req.CitizenId,
		req.RequestType,
		req.Description,
		req.Status,
	).Scan(
		&serviceReq.RequestId,
		&serviceReq.CitizenId,
		&serviceReq.RequestType,
		&serviceReq.Description,
		&serviceReq.Status,
		&serviceReq.CreateAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error creating service request")
		slog.Error("Error creating service request", err)
		return nil, err
	}

	return &serviceReq, nil
}

// GetCitizenRequests retrieves all service requests for a citizen by their ID
func (r *ServiceRequest) GetCitizenRequests(ctx context.Context, req *citi.ById) (*citi.ServiceReqList, error) {
	query := `
    SELECT 
        sr.RequestId,
        c.CitizenId,
        c.FirstName,
        c.LastName,
        c.DateOfBirth,
        c.Address,
        c.PhoneNumber,
        c.Email,
        sr.RequestType,
        sr.Description,
        sr.Status,
        sr.CreateAt,
        sr.UpdateAt
    FROM service_requests sr
    INNER JOIN citizens c ON sr.CitizenId = c.CitizenId
    WHERE sr.CitizenId = $1 AND sr.DeleteAt IS NULL`

	rows, err := r.db.Query(ctx, query, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Error retrieving service requests")
		slog.Error("Error retrieving service requests", err)
		return nil, err
	}
	defer rows.Close()

	var serviceReqs []*citi.ServiceReq
	for rows.Next() {
		serviceReq := &citi.ServiceReq{}
		citizenModel := &citi.CitizenModel{}
		err := rows.Scan(
			&serviceReq.RequestId,
			&citizenModel.UserId,
			&citizenModel.FirstName,
			&citizenModel.LastName,
			&citizenModel.DateOfBirth,
			&citizenModel.Address,
			&citizenModel.PhoneNumber,
			&citizenModel.Email,
			&serviceReq.RequestType,
			&serviceReq.Description,
			&serviceReq.Status,
			&serviceReq.CreateAt,
			&serviceReq.UpdateAt,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error scanning service request rows")
			slog.Error("Error scanning service request rows", err)
			return nil, err
		}
		serviceReq.CitizenId = citizenModel
		serviceReqs = append(serviceReqs, serviceReq)
	}

	if err := rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error iterating over service request rows")
		slog.Error("Error iterating over service request rows", err)
		return nil, err
	}

	return &citi.ServiceReqList{
		Model: serviceReqs,
		Count: int32(len(serviceReqs)),
	}, nil
}


// UpdateServiceRequest updates an existing service request in the database
func (r *ServiceRequest) UpdateServiceRequest(ctx context.Context, req *citi.ServiceReqCreate) (*citi.ServiceReq, error) {
	query := `
    UPDATE service_requests
    SET 
        request_type = $2,
        description = $3,
        status = $4,
        update_at = NOW()
    WHERE request_id = $1 AND delete_at = 0
    RETURNING 
        RequestId,
        CitizenId,
        RequestType,
        Description,
        Status,
        CreateAt,
        UpdateAt`

	serviceReq := citi.ServiceReq{}
	err := r.db.QueryRow(ctx, query,
		req.RequestId,
		req.RequestType,
		req.Description,
		req.Status,
	).Scan(
		&serviceReq.RequestId,
		&serviceReq.CitizenId,
		&serviceReq.RequestType,
		&serviceReq.Description,
		&serviceReq.Status,
		&serviceReq.CreateAt,
		&serviceReq.UpdateAt,
	)

	if err != nil {
		log.Error().Err(err).Msg("Error updating service request")
		slog.Error("Error updating service request", err)
		return nil, err
	}

	return &serviceReq, nil
}

// DeleteServiceRequest performs a soft delete on a service request by its ID
func (r *ServiceRequest) DeleteServiceRequest(ctx context.Context, req *citi.ById) (*citi.Void, error) {
	query := `
    UPDATE service_requests
    SET delete_at = NOW()
    WHERE request_id = $1 AND delete_at = 0`

	res, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Error().Err(err).Msg("Error performing soft delete on service request")
		slog.Error("Error performing soft delete on service request", err)
		return nil, err
	}

	// Check if the update affected any rows
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		err = sql.ErrNoRows
		log.Warn().Err(err).Msg("No service request found or already soft deleted")
		slog.Warn("No service request found or already soft deleted", err)
		return nil, err
	}

	return &citi.Void{}, nil
}
