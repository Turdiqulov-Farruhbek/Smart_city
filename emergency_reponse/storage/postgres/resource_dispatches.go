package postgres

import (
	"database/sql"
	er "emergency_response-service/genproto/emergency_response"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type ResourceDispatchesRepo struct {
	db *sql.DB
}

func NewResourceDispatchesRepo(db *sql.DB) *ResourceDispatchesRepo {
	return &ResourceDispatchesRepo{db: db}
}

func (rd *ResourceDispatchesRepo) Create(req *er.DispatchesCreateReq) (*er.DispatchesRes, error) {
	id := uuid.New().String()
	dispatch := er.DispatchesRes{}

	query := `
	INSERT INTO resource_dispatches(
		id,
		incident_id,
		resource_id,
		dispatched_at,
		arrived_at
	) VALUES ($1, $2, $3, $4, $5)
	RETURNING 
		id,
		incident_id,
		resource_id,
		dispatched_at,
		arrived_at
	`

	row := rd.db.QueryRow(query, id, req.IncidentId, req.ResourceId, req.DispatchedAt, req.ArrivedAt)

	err := row.Scan(
		&dispatch.Id,
		&dispatch.IncidentId,
		&dispatch.ResourceId,
		&dispatch.DispatchedAt,
		&dispatch.ArrivedAt,
	)

	if err != nil {
		log.Println("Error while creating dispatches: ", err)
		return nil, err
	}

	log.Println("Successfully created dispatch")

	return &dispatch, nil
}

func (rd *ResourceDispatchesRepo) GetById(id *er.ById) (*er.DispatchesFullRes, error) {
	dispatches := er.DispatchesFullRes{
		Incident: &er.IncidentsRes{},
		Resource: &er.ResourcesRes{},
	}

	query := `
	SELECT 
		d.id,
		d.dispatched_at,
		d.arrived_at,
		i.id as incident_id,
		i.type,
		i.location,
		i.description,
		i.status,
		i.reporter_at,
		r.id as resource_id,
		r.type,
		r.current_location,
		r.status
	FROM 
		resource_dispatches as d
	JOIN 
		emergency_incident as i ON i.id = d.incident_id AND i.deleted_at = 0
	JOIN 
		emergency_resource as r ON r.id = d.resource_id AND r.deleted_at = 0
	WHERE 
		d.id = $1
	AND 
		d.deleted_at = 0	
	`

	row := rd.db.QueryRow(query, id.Id)

	err := row.Scan(
		&dispatches.Id,
		&dispatches.DispatchedAt,
		&dispatches.ArrivedAt,
		&dispatches.Incident.Id,
		&dispatches.Incident.Type,
		&dispatches.Incident.Location,
		&dispatches.Incident.Description,
		&dispatches.Incident.Status,
		&dispatches.Incident.ReportedAt,
		&dispatches.Resource.Id,
		&dispatches.Resource.Type,
		&dispatches.Resource.CurrentLocation,
		&dispatches.Resource.Status,
	)

	if err != nil {
		log.Println("Error while getting dispatch by id: ", err)
		return nil, err
	}

	log.Println("Successfully got dispatch")

	return &dispatches, nil
}

func (rd *ResourceDispatchesRepo) GetAll(filter *er.Filter) (*er.DispatchesGetAllRes, error) {
	dispatches := er.DispatchesGetAllRes{
		Dispatches: []*er.DispatchesFullRes{},
	}

	query := `
	SELECT 
		d.id,
		d.dispatched_at,
		d.arrived_at,
		i.id as incident_id,
		i.type,
		i.location,
		i.description,
		i.status,
		i.reporter_at,
		r.id as resource_id,
		r.type,
		r.current_location,
		r.status
	FROM 
		resource_dispatches as d
	JOIN 
		emergency_incident as i ON i.id = d.incident_id AND i.deleted_at = 0
	JOIN 
		emergency_resource as r ON r.id = d.resource_id AND r.deleted_at = 0
	WHERE 
		d.deleted_at = 0	
	`

	var args []interface{}

	var defaultLimit int32
	err := rd.db.QueryRow("SELECT COUNT(1) FROM resource_dispatches WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := rd.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving dispatches: ", err)
		return nil, err
	}

	for rows.Next() {
		dispatch := er.DispatchesFullRes{
			Incident: &er.IncidentsRes{},
			Resource: &er.ResourcesRes{},
		}

		err := rows.Scan(
			&dispatch.Id,
			&dispatch.DispatchedAt,
			&dispatch.ArrivedAt,
			&dispatch.Incident.Id,
			&dispatch.Incident.Type,
			&dispatch.Incident.Location,
			&dispatch.Incident.Description,
			&dispatch.Incident.Status,
			&dispatch.Incident.ReportedAt,
			&dispatch.Resource.Id,
			&dispatch.Resource.Type,
			&dispatch.Resource.CurrentLocation,
			&dispatch.Resource.Status,
		)

		if err != nil {
			log.Println("Error while scanning all ditpatches: ", err)
			return nil, err
		}

		dispatches.Dispatches = append(dispatches.Dispatches, &dispatch)
	}

	log.Println("Successfully fetched all dispatches")

	return &dispatches, nil
}

func (rd *ResourceDispatchesRepo) Update(req *er.DispatchesUpdateReq) (*er.DispatchesRes, error) {
	dispatch := er.DispatchesRes{}

	query := `
	UPDATE
		resource_dispatches
	SET 
		updated_at = now()
	`

	var conditions []string
	var args []interface{}

	if req.Dispatch.IncidentId != "" && req.Dispatch.IncidentId != "string" {
		conditions = append(conditions, " incident_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Dispatch.IncidentId)
	}
	if req.Dispatch.ResourceId != "" && req.Dispatch.ResourceId != "string" {
		conditions = append(conditions, " resource_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Dispatch.ResourceId)
	}
	if req.Dispatch.DispatchedAt != "" && req.Dispatch.DispatchedAt != "string" {
		conditions = append(conditions, " dispatched_at = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Dispatch.DispatchedAt)
	}
	if req.Dispatch.ArrivedAt != "" && req.Dispatch.ArrivedAt != "string" {
		conditions = append(conditions, " arrived_at = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Dispatch.ArrivedAt)
	}
	if len(conditions) > 0 {
		query += strings.Join(conditions, ", ")
	}

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	query += `
	RETURNING
		id,
		incident_id,
		resource_id,
		dispatched_at,
		arrived_at
	`
	args = append(args, req.Id)

	row := rd.db.QueryRow(query, args...)

	err := row.Scan(
		&dispatch.Id,
		&dispatch.IncidentId,
		&dispatch.ResourceId,
		&dispatch.DispatchedAt,
		&dispatch.ArrivedAt,
	)

	if err != nil {
		log.Println("Error while updating dispatch: ", err)
		return nil, err
	}

	log.Println("Successfully updated dispatch")

	return &dispatch, nil
}

func (rd *ResourceDispatchesRepo) Delete(id *er.ById) (*er.Void, error) {
	void := er.Void{}

	query := `
	UPDATE 
		resource_dispatches
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := rd.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting dispatch: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("dispatch with id %s not found", id.Id)
	}

	log.Println("Successfully deleted dispatch")

	return &void, nil
}
