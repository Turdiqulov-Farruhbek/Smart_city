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

type EmergencyDrillsRepo struct {
	db *sql.DB
}

func NewEmergencyDrillsRepo(db *sql.DB) *EmergencyDrillsRepo {
	return &EmergencyDrillsRepo{db: db}
}

func (ed *EmergencyDrillsRepo) Create(req *er.DrillsCreateReq) (*er.DrillsRes, error) {
	id := uuid.New().String()
	drill := er.DrillsRes{}

	query := `
	INSERT INTO emergancy_alerts(
		id,
		type,
		location,
		scheduled_at,
		completed_at
	) VALUES ($1, $2, $3, $4, $5)
	RETURNING 
		id,
		type,
		location,
		scheduled_at,
		completed_at
	`

	row := ed.db.QueryRow(query, id, req.Type, req.Location, req.ScheduledAt, req.CompletedAt)

	err := row.Scan(
		&drill.Id,
		&drill.Type,
		&drill.Location,
		&drill.ScheduledAt,
		&drill.CompletedAt,
	)

	if err != nil {
		log.Println("Error while creating drills: ", err)
		return nil, err
	}

	log.Println("Successfully created drill")

	return &drill, nil
}

func (ed *EmergencyDrillsRepo) GetById(id *er.ById) (*er.DrillsGetByIdRes, error) {
	drill := er.DrillsGetByIdRes{}

	query := `
	SELECT 
		id,
		type,
		location,
		scheduled_at,
		completed_at
	FROM 
		emergency_drills
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := ed.db.QueryRow(query, id.Id)

	err := row.Scan(
		&drill.Drill.Id,
		&drill.Drill.Type,
		&drill.Drill.Location,
		&drill.Drill.ScheduledAt,
		&drill.Drill.CompletedAt,
	)

	if err != nil {
		log.Println("Error while getting drill by id: ", err)
		return nil, err
	}

	log.Println("Successfully got drill")

	return &drill, nil
}

func (ed *EmergencyDrillsRepo) GetAll(filter *er.Filter) (*er.DrillsGetAllRes, error) {
	drills := er.DrillsGetAllRes{}

	query := `
	SELECT 
		id,
		type,
		location,
		scheduled_at,
		completed_at
	FROM 
		emergency_drills
	WHERE 
		deleted_at = 0
	`

	var args []interface{}

	var defaultLimit int32
	err := ed.db.QueryRow("SELECT COUNT(1) FROM emergency_drills WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := ed.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving drills: ", err)
		return nil, err
	}

	for rows.Next() {
		drill := er.DrillsRes{}

		err := rows.Scan(
			&drill.Id,
			&drill.Type,
			&drill.Location,
			&drill.ScheduledAt,
			&drill.CompletedAt,
		)

		if err != nil {
			log.Println("Error while scanning all drills: ", err)
			return nil, err
		}

		drills.Drills = append(drills.Drills, &drill)
	}

	log.Println("Successfully fetched all drills")

	return &drills, nil
}

func (ed *EmergencyDrillsRepo) Update(req *er.DrillsUpdateReq) (*er.DrillsRes, error) {
	drill := er.DrillsRes{}

	query := `
	UPDATE
		emergency_drills
	SET 
		updated_at = now()
	`

	var conditions []string
	var args []interface{}

	if req.Drill.Type != "" && req.Drill.Type != "string" {
		conditions = append(conditions, " type = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Drill.Type)
	}
	if req.Drill.Location != "" && req.Drill.Location != "string" {
		conditions = append(conditions, " location = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Drill.Location)
	}
	if req.Drill.ScheduledAt != "" && req.Drill.ScheduledAt != "string" {
		conditions = append(conditions, " scheduled_at = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Drill.ScheduledAt)
	}
	if req.Drill.CompletedAt != "" && req.Drill.CompletedAt != "string" {
		conditions = append(conditions, " completed_at = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Drill.CompletedAt)
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
		scheduled_at,
		completed_at
	`
	args = append(args, req.Id)

	row := ed.db.QueryRow(query, args...)

	err := row.Scan(
		&drill.Id,
		&drill.Type,
		&drill.Location,
		&drill.ScheduledAt,
		&drill.CompletedAt,
	)

	if err != nil {
		log.Println("Error while updating drill: ", err)
		return nil, err
	}

	log.Println("Successfully updated drill")

	return &drill, nil
}

func (ed *EmergencyDrillsRepo) Delete(id *er.ById) (*er.Void, error) {
	void := er.Void{}

	query := `
	UPDATE 
		emergency_drills
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := ed.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting drill: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("drill with id %s not found", id.Id)
	}

	log.Println("Successfully deleted drill")

	return &void, nil
}
