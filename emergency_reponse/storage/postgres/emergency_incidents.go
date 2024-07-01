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

type EmergencyIncidentsRepo struct {
	db *sql.DB
}

func NewEmergencyIncidentsRepo(db *sql.DB) *EmergencyIncidentsRepo {
	return &EmergencyIncidentsRepo{db: db}
}

func (ei *EmergencyIncidentsRepo) Create(req *er.IncidentsCreateReq) (*er.IncidentsRes, error) {
	id := uuid.New().String()
	incident := er.IncidentsRes{}

	query := `
	INSERT INTO emergancy_incidents(
		id,
		type,
		location,
		description,
		status,
		reporter_at
	) VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING 
		id,
		type,
		location,
		description,
		status,
		reporter_at		
	`

	row := ei.db.QueryRow(query, id, req.Type, req.Location, req.Description, req.Status, req.ReportedAt)

	err := row.Scan(
		&incident.Id,
		&incident.Type,
		&incident.Location,
		&incident.Description,
		&incident.Status,
		&incident.ReportedAt,
	)

	if err != nil {
		log.Println("Error while creating incidents: ", err)
		return nil, err
	}

	log.Println("Successfully created incident")

	return &incident, nil
}

func (ei *EmergencyIncidentsRepo) GetById(id *er.ById) (*er.IncidentsGetByIdRes, error) {

	incident := er.IncidentsGetByIdRes{}

	query := `
	SELECT 
		id,
		type,
		location,
		description,
		status,
		reporter_at
	FROM 
		emergency_incidents
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := ei.db.QueryRow(query, id.Id)

	err := row.Scan(
		&incident.Incident.Id,
		&incident.Incident.Type,
		&incident.Incident.Location,
		&incident.Incident.Description,
		&incident.Incident.Status,
		&incident.Incident.ReportedAt,
	)

	if err != nil {
		log.Println("Error while getting incident by id: ", err)
		return nil, err
	}

	log.Println("Successfully got incident")

	return &incident, nil
}

func (ei *EmergencyIncidentsRepo) GetAll(filter *er.Filter) (*er.IncidentsGetAllRes, error) {

	incidents := er.IncidentsGetAllRes{}

	query := `
	SELECT 
		id,
		type,
		location,
		description,
		status,
		reporter_at
	FROM 
		emergency_incidents
	WHERE 
		deleted_at = 0	
	`

	var args []interface{}

	var defaultLimit int32
	err := ei.db.QueryRow("SELECT COUNT(1) FROM emergency_incidents WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := ei.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving incidents: ", err)
		return nil, err
	}

	for rows.Next() {
		incident := er.IncidentsRes{}

		err := rows.Scan(
			&incident.Id,
			&incident.Type,
			&incident.Location,
			&incident.Description,
			&incident.Status,
			&incident.ReportedAt,
		)

		if err != nil {
			log.Println("Error while scanning all incidents: ", err)
			return nil, err
		}

		incidents.Incidents = append(incidents.Incidents, &incident)
	}

	log.Println("Successfully fetched all incidents")

	return &incidents, nil
}

func (ei *EmergencyIncidentsRepo) Update(req *er.IncidentsUpdateReq) (*er.IncidentsRes, error) {

	incident := er.IncidentsRes{}

	query := `
	UPDATE
		emergency_incidents
	SET 
		updated_at = now()
	`

	var conditions []string
	var args []interface{}

	if req.Incident.Type != "" && req.Incident.Type != "string" {
		conditions = append(conditions, " type = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Incident.Type)
	}
	if req.Incident.Location != "" && req.Incident.Location != "string" {
		conditions = append(conditions, " location = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Incident.Location)
	}
	if req.Incident.Description != "" && req.Incident.Description != "string" {
		conditions = append(conditions, " description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Incident.Description)
	}
	if req.Incident.Status != "" && req.Incident.Status != "string" {
		conditions = append(conditions, " status = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Incident.Status)
	}
	if req.Incident.ReportedAt != "" && req.Incident.ReportedAt != "string" {
		conditions = append(conditions, " reporter_at = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Incident.ReportedAt)
	}
	if len(conditions) > 0 {
		query += strings.Join(conditions, ", ")
	}

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	query += `
	RETURNING
		id,
		type,
		location,
		description,
		status,
		reporter_at
	`
	args = append(args, req.Id)

	row := ei.db.QueryRow(query, args...)

	err := row.Scan(
		&incident.Id,
		&incident.Type,
		&incident.Location,
		&incident.Description,
		&incident.Status,
		&incident.ReportedAt,
	)

	if err != nil {
		log.Println("Error while updating incident: ", err)
		return nil, err
	}

	log.Println("Successfully updated incident")

	return &incident, nil
}
func (ei *EmergencyIncidentsRepo) Delete(id *er.ById) (*er.Void, error) {

	void := er.Void{}

	query := `
	UPDATE 
		emergency_incidents
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := ei.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting incident: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("incident with id %s not found", id.Id)
	}

	log.Println("Successfully deleted incident")

	return &void, nil
}

func (ei *EmergencyIncidentsRepo) GetActiveIncidents(filter *er.Filter)(*er.IncidentsGetAllRes, error){
	incidents := er.IncidentsGetAllRes{}

	query := `
	SELECT 
		id,
		type,
		location,
		description,
		status,
		reporter_at
	FROM 
		emergency_incidents
	WHERE 
		deleted_at = 0	
	AND 
		status = 'active'
	`

	var args []interface{}

	var defaultLimit int32
	err := ei.db.QueryRow("SELECT COUNT(1) FROM emergency_incidents WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := ei.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving incidents: ", err)
		return nil, err
	}

	for rows.Next() {
		incident := er.IncidentsRes{}

		err := rows.Scan(
			&incident.Id,
			&incident.Type,
			&incident.Location,
			&incident.Description,
			&incident.Status,
			&incident.ReportedAt,
		)

		if err != nil {
			log.Println("Error while scanning all incidents: ", err)
			return nil, err
		}

		incidents.Incidents = append(incidents.Incidents, &incident)
	}

	log.Println("Successfully fetched all incidents")

	return &incidents, nil
}