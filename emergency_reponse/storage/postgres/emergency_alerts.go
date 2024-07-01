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

type EmergencyAlertsRepo struct {
	db *sql.DB
}

func NewEmergencyAlertsRepo(db *sql.DB) *EmergencyAlertsRepo {
	return &EmergencyAlertsRepo{db: db}
}

func (ea *EmergencyAlertsRepo) Create(req *er.AlertsCreateReq) (*er.AlertsRes, error) {
	id := uuid.New().String()
	alert := er.AlertsRes{}

	query := `
	INSERT INTO emergancy_alerts(
		id,
		type,
		message,
		affected_area,
		issued_at
	) VALUES ($1, $2, $3, $4, $5)
	RETURNING 
		id,
		type,
		message,
		affected_area,
		issued_at
	`

	row := ea.db.QueryRow(query, id, req.Type, req.Message, req.AffectedArea, req.IssuedAt)

	err := row.Scan(
		&alert.Id,
		&alert.Type,
		&alert.Message,
		&alert.AffectedArea,
		&alert.IssuedAt,
	)

	if err != nil {
		log.Println("Error while creating alerts: ", err)
		return nil, err
	}

	log.Println("Successfully created alert")

	return &alert, nil
}

func (ea *EmergencyAlertsRepo) GetById(id *er.ById) (*er.AlertsGetByIdRes, error) {
	alert := er.AlertsGetByIdRes{}

	query := `
	SELECT 
		id,
		type,
		message,
		affected_area,
		issued_at
	FROM 
		emergency_alerts
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := ea.db.QueryRow(query, id.Id)

	err := row.Scan(
		&alert.Alert.Id,
		&alert.Alert.Type,
		&alert.Alert.Message,
		&alert.Alert.AffectedArea,
		&alert.Alert.IssuedAt,
	)

	if err != nil {
		log.Println("Error while getting alert by id: ", err)
		return nil, err
	}

	log.Println("Successfully got alert")

	return &alert, nil
}

func (ea *EmergencyAlertsRepo) GetAll(filter *er.Filter) (*er.AlertsGetAllRes, error) {

	alerts := er.AlertsGetAllRes{}

	query := `
	SELECT 
		id,
		type,
		message,
		affected_area,
		issued_at
	FROM 
		emergency_alerts
	WHERE 
		deleted_at = 0
	`

	var args []interface{}

	var defaultLimit int32
	err := ea.db.QueryRow("SELECT COUNT(1) FROM emergency_alerts WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := ea.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving alerts: ", err)
		return nil, err
	}

	for rows.Next() {
		alert := er.AlertsRes{}

		err := rows.Scan(
			&alert.Id,
			&alert.Type,
			&alert.Message,
			&alert.AffectedArea,
			&alert.IssuedAt,
		)

		if err != nil {
			log.Println("Error while scanning all aletrs: ", err)
			return nil, err
		}

		alerts.Alerts = append(alerts.Alerts, &alert)
	}

	log.Println("Successfully fetched all alerts")

	return &alerts, nil
}

func (ea *EmergencyAlertsRepo) Update(req *er.AlertsUpdateReq) (*er.AlertsRes, error) {
	alert := er.AlertsRes{}

	query := `
	UPDATE
		emergency_alerts
	SET 
		updated_at = now()
	`

	var conditions []string
	var args []interface{}

	if req.Alert.Type != "" && req.Alert.Type != "string" {
		conditions = append(conditions, " type = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Alert.Type)
	}
	if req.Alert.Message != "" && req.Alert.Message != "string" {
		conditions = append(conditions, " message = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Alert.Message)
	}
	if req.Alert.AffectedArea != "" && req.Alert.AffectedArea != "string" {
		conditions = append(conditions, " affected_area = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Alert.AffectedArea)
	}
	if req.Alert.IssuedAt != "" && req.Alert.IssuedAt != "string" {
		conditions = append(conditions, " issued_at = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Alert.IssuedAt)
	}
	if len(conditions) > 0 {
		query += strings.Join(conditions, ", ")
	}

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	query += `
	RETURNING
		id,
		type,
		message,
		affected_area,
		issued_at
	`
	args = append(args, req.Id)

	row := ea.db.QueryRow(query, args...)

	err := row.Scan(
		&alert.Id,
		&alert.Type,
		&alert.Message,
		&alert.AffectedArea,
		&alert.IssuedAt,
	)

	if err != nil {
		log.Println("Error while updating alert: ", err)
		return nil, err
	}

	log.Println("Successfully updated alert")

	return &alert, nil
}

func (ea *EmergencyAlertsRepo) Delete(id *er.ById) (*er.Void, error) {
	void := er.Void{}

	query := `
	UPDATE 
		emergency_alerts
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := ea.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting alert: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("alert with id %s not found", id.Id)
	}

	log.Println("Successfully deleted alert")

	return &void, nil
}
