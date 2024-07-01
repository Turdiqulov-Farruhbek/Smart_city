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

type EmergencyFacilitiesRepo struct {
	db *sql.DB
}

func NewEmergencyFacilitiesRepo(db *sql.DB) *EmergencyFacilitiesRepo {
	return &EmergencyFacilitiesRepo{db: db}
}

func (ef *EmergencyFacilitiesRepo) Create(req *er.FacilitiesCreateReq) (*er.FacilitiesRes, error) {
	id := uuid.New().String()
	facility := er.FacilitiesRes{}

	query := `
	INSERT INTO emergancy_facilities(
		id,
		type,
		name,
		address,
		capacity
	) VALUES ($1, $2, $3, $4, $5)
	RETURNING 
		id,
		type,
		name,
		address,
		capacity
	`

	row := ef.db.QueryRow(query, id, req.Type, req.Name, req.Address, req.Capacity)

	err := row.Scan(
		&facility.Id,
		&facility.Type,
		&facility.Name,
		&facility.Address,
		&facility.Capacity,
	)

	if err != nil {
		log.Println("Error while creating facilities: ", err)
		return nil, err
	}

	log.Println("Successfully created facility")

	return &facility, nil
}

func (ef *EmergencyFacilitiesRepo) GetById(id *er.ById) (*er.FacilitiesGetByIdRes, error) {
	facility := er.FacilitiesGetByIdRes{}

	query := `
	SELECT 
		id,
		type,
		name,
		address,
		capacity
	FROM 
		emergency_facilities
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := ef.db.QueryRow(query, id.Id)

	err := row.Scan(
		&facility.Facelity.Id,
		&facility.Facelity.Type,
		&facility.Facelity.Name,
		&facility.Facelity.Address,
		&facility.Facelity.Capacity,
	)

	if err != nil {
		log.Println("Error while getting facility by id: ", err)
		return nil, err
	}

	log.Println("Successfully got facility")

	return &facility, nil
}

func (ef *EmergencyFacilitiesRepo) GetAll(filter *er.Filter) (*er.FacilitiesGetAllRes, error) {
	facilities := er.FacilitiesGetAllRes{}

	query := `
	SELECT 
		id,
		type,
		name,
		address,
		capacity
	FROM 
		emergency_facilities
	WHERE 
		deleted_at = 0	
	`

	var args []interface{}

	var defaultLimit int32
	err := ef.db.QueryRow("SELECT COUNT(1) FROM emergency_facilities WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := ef.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving facilities: ", err)
		return nil, err
	}

	for rows.Next() {
		facility := er.FacilitiesRes{}

		err := rows.Scan(
			&facility.Id,
			&facility.Type,
			&facility.Name,
			&facility.Address,
			&facility.Capacity,
		)

		if err != nil {
			log.Println("Error while scanning all facilites: ", err)
			return nil, err
		}

		facilities.Facilities = append(facilities.Facilities, &facility)
	}

	log.Println("Successfully fetched all facilities")

	return &facilities, nil
}

func (ef *EmergencyFacilitiesRepo) Update(req *er.FacilitiesUpdateReq) (*er.FacilitiesRes, error) {
	facility := er.FacilitiesRes{}

	query := `
	UPDATE
		emergency_facilities
	SET 
		updated_at = now()
	`

	var conditions []string
	var args []interface{}

	if req.Facelity.Type != "" && req.Facelity.Type != "string" {
		conditions = append(conditions, " type = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Facelity.Type)
	}
	if req.Facelity.Name != "" && req.Facelity.Name != "string" {
		conditions = append(conditions, " name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Facelity.Name)
	}
	if req.Facelity.Address != "" && req.Facelity.Address != "string" {
		conditions = append(conditions, " address = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Facelity.Address)
	}
	if req.Facelity.Capacity != 0 {
		conditions = append(conditions, " capacity = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Facelity.Capacity)
	}
	if len(conditions) > 0 {
		query += strings.Join(conditions, ", ")
	}

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	query += `
	RETURNING
		id,
		type,
		name,
		address,
		capacity
	`
	args = append(args, req.Id)

	row := ef.db.QueryRow(query, args...)

	err := row.Scan(
		&facility.Id,
		&facility.Type,
		&facility.Name,
		&facility.Address,
		&facility.Capacity,
	)

	if err != nil {
		log.Println("Error while updating facility: ", err)
		return nil, err
	}

	log.Println("Successfully updated facility")

	return &facility, nil
}

func (ef *EmergencyFacilitiesRepo) Delete(id *er.ById) (*er.Void, error) {
	void := er.Void{}

	query := `
	UPDATE 
		emergency_facilities
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := ef.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting facility: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("facelity with id %s not found", id.Id)
	}

	log.Println("Successfully deleted facelity")

	return &void, nil
}

func (ef *EmergencyFacilitiesRepo) GetFacilityByType(filter *er.FilterType)(*er.FacilitiesGetAllRes, error){
	facilities := er.FacilitiesGetAllRes{}

	query := `
	SELECT 
		id,
		type,
		name,
		address,
		capacity
	FROM 
		emergency_facilities
	WHERE 
		deleted_at = 0	
	`

	var args []interface{}
	var conditions []string

	if filter.Type != "" && filter.Type != "string"{
		conditions = append(conditions, " type = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Type)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var defaultLimit int32
	err := ef.db.QueryRow("SELECT COUNT(1) FROM emergency_facilities WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := ef.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving facilities: ", err)
		return nil, err
	}

	for rows.Next() {
		facility := er.FacilitiesRes{}

		err := rows.Scan(
			&facility.Id,
			&facility.Type,
			&facility.Name,
			&facility.Address,
			&facility.Capacity,
		)

		if err != nil {
			log.Println("Error while scanning all facilites: ", err)
			return nil, err
		}

		facilities.Facilities = append(facilities.Facilities, &facility)
	}

	log.Println("Successfully fetched all facilities")

	return &facilities, nil
}
