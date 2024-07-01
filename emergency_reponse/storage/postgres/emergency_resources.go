package postgres

import (
	"database/sql"
	ers "emergency_response-service/genproto/emergency_response"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type EmergencyResourcesRepo struct {
	db *sql.DB
}

func NewEmergencyResourcesRepo(db *sql.DB) *EmergencyResourcesRepo {
	return &EmergencyResourcesRepo{db: db}
}

func (er *EmergencyResourcesRepo) Create(req *ers.ResourcesCreateReq) (*ers.ResourcesRes, error) {
	id := uuid.New().String()
	resource := ers.ResourcesRes{}

	query := `
	INSERT INTO emergancy_resources(
		id,
		type,
		current_location,
		status
	) VALUES ($1, $2, $3, $4)
	RETURNING 
		id,
		type,
		current_location,
		status		
	`

	row := er.db.QueryRow(query, id, req.Type, req.CurrentLocation, req.Status)

	err := row.Scan(
		&resource.Id,
		&resource.Type,
		&resource.CurrentLocation,
		&resource.Status,
	)

	if err != nil {
		log.Println("Error while creating resources: ", err)
		return nil, err
	}

	log.Println("Successfully created resource")

	return &resource, nil
}

func (er *EmergencyResourcesRepo) GetById(id *ers.ById) (*ers.ResourcesGetByIdRes, error) {
	resource := ers.ResourcesGetByIdRes{}

	query := `
	SELECT 
		id,
		type,
		current_location,
		status
	FROM 
		emergency_resources
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := er.db.QueryRow(query, id.Id)

	err := row.Scan(
		&resource.Resource.Id,
		&resource.Resource.Type,
		&resource.Resource.CurrentLocation,
		&resource.Resource.Status,
	)

	if err != nil {
		log.Println("Error while getting resource by id: ", err)
		return nil, err
	}

	log.Println("Successfully got resource")

	return &resource, nil
}

func (er *EmergencyResourcesRepo) GetAll(filter *ers.Filter) (*ers.ResourcesGetAllRes, error) {
	resources := ers.ResourcesGetAllRes{}

	query := `
	SELECT 
		id,
		type,
		current_location,
		status
	FROM 
		emergency_resources
	WHERE 
		deleted_at = 0
	`

	var args []interface{}

	var defaultLimit int32
	err := er.db.QueryRow("SELECT COUNT(1) FROM emergency_resources WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := er.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving resources: ", err)
		return nil, err
	}

	for rows.Next() {
		resource := ers.ResourcesRes{}

		err := rows.Scan(
			&resource.Id,
			&resource.Type,
			&resource.CurrentLocation,
			&resource.Status,
		)

		if err != nil {
			log.Println("Error while scanning all resources: ", err)
			return nil, err
		}

		resources.Resources = append(resources.Resources, &resource)
	}

	log.Println("Successfully fetched all resources")

	return &resources, nil

}

func (er *EmergencyResourcesRepo) Update(req *ers.ResourcesUpdateReq) (*ers.ResourcesRes, error) {
	resource := ers.ResourcesRes{}

	query := `
	UPDATE
		emergency_resources
	SET 
		updated_at = now()
	`

	var conditions []string
	var args []interface{}

	if req.Resource.Type != "" && req.Resource.Type != "string" {
		conditions = append(conditions, " type = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Resource.Type)
	}
	if req.Resource.CurrentLocation != "" && req.Resource.CurrentLocation != "string" {
		conditions = append(conditions, " current_location = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Resource.CurrentLocation)
	}
	if req.Resource.Status != "" && req.Resource.Status != "string" {
		conditions = append(conditions, " status = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Resource.Status)
	}
	if len(conditions) > 0 {
		query += strings.Join(conditions, ", ")
	}

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	query += `
	RETURNING
		id,
		type,
		current_location,
		status
	`
	args = append(args, req.Id)

	row := er.db.QueryRow(query, args...)

	err := row.Scan(
		&resource.Id,
		&resource.Type,
		&resource.CurrentLocation,
		&resource.Status,
	)

	if err != nil {
		log.Println("Error while updating incident: ", err)
		return nil, err
	}

	log.Println("Successfully updated incident")

	return &resource, nil
}

func (er *EmergencyResourcesRepo) Delete(id *ers.ById) (*ers.Void, error) {
	void := ers.Void{}

	query := `
	UPDATE 
		emergency_resources
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := er.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting resource: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("resource with id %s not found", id.Id)
	}

	log.Println("Successfully deleted resource")

	return &void, nil
}

func (er *EmergencyResourcesRepo) GetAvailableResources(filter *ers.FilterType)(*ers.ResourcesGetAllRes, error){
	resources := ers.ResourcesGetAllRes{}

	query := `
	SELECT 
		id,
		type,
		current_location,
		status
	FROM 
		emergency_resources
	WHERE 
		deleted_at = 0
	AND
		status = 'available'
	`
	var args []interface{}
	var conditions []string

	if filter.Type != "" && filter.Type != "string"{
		conditions = append(conditions, " type = $"+strconv.Itoa(len(args)+1))
		args = append(args, filter.Type)
	}

	var defaultLimit int32
	err := er.db.QueryRow("SELECT COUNT(1) FROM emergency_resources WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := er.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving resources: ", err)
		return nil, err
	}

	for rows.Next() {
		resource := ers.ResourcesRes{}

		err := rows.Scan(
			&resource.Id,
			&resource.Type,
			&resource.CurrentLocation,
			&resource.Status,
		)

		if err != nil {
			log.Println("Error while scanning all resources: ", err)
			return nil, err
		}

		resources.Resources = append(resources.Resources, &resource)
	}

	log.Println("Successfully fetched all resources")

	return &resources, nil
}