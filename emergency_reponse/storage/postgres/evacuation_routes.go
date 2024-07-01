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

type EvacuationRoutesRepo struct {
	db *sql.DB
}

func NewEvacuationRoutesRepo(db *sql.DB) *EvacuationRoutesRepo {
	return &EvacuationRoutesRepo{db: db}
}

func (evr *EvacuationRoutesRepo) Create(req *er.RoutesCreateReq) (*er.RoutesRes, error) {
	id := uuid.New().String()
	route := er.RoutesRes{}

	query := `
	INSERT INTO evacuation_routes(
		id,
		start_point,
		end_point,
		description
	) VALUES ($1, $2, $3, $4)
	RETURNING 
		id,
		start_point,
		end_point,
		description	
	`

	row := evr.db.QueryRow(query, id, req.StartPoint, req.EndPoint, req.Description)

	err := row.Scan(
		&route.Id,
		&route.StartPoint,
		&route.EndPoint,
		&route.Description,
	)

	if err != nil {
		log.Println("Error while creating routes: ", err)
		return nil, err
	}

	log.Println("Successfully created routes")

	return &route, nil
}

func (evr *EvacuationRoutesRepo) GetById(id *er.ById) (*er.RoutesGetByIdRes, error) {
	route := er.RoutesGetByIdRes{}

	query := `
	SELECT 
		id,
		start_point,
		end_point,
		description	
	FROM 
		evacuation_routes
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := evr.db.QueryRow(query, id.Id)

	err := row.Scan(
		&route.Route.Id,
		&route.Route.StartPoint,
		&route.Route.EndPoint,
		&route.Route.Description,
	)

	if err != nil {
		log.Println("Error while getting route by id: ", err)
		return nil, err
	}

	log.Println("Successfully got route")

	return &route, nil
}

func (evr *EvacuationRoutesRepo) GetAll(filter *er.Filter) (*er.RoutesGetAllRes, error) {
	routes := er.RoutesGetAllRes{}

	query := `
	SELECT 
		id,
		start_point,
		end_point,
		description	
	FROM 
		evacuation_routes
	WHERE 
		deleted_at = 0
	`

	var args []interface{}

	var defaultLimit int32
	err := evr.db.QueryRow("SELECT COUNT(1) FROM evacuation_routes WHERE deleted_at=0").Scan(&defaultLimit)
	if err != nil {
		log.Println("Error while getting count : ", err)
		return nil, err
	}
	if filter.Limit == 0 {
		filter.Limit = defaultLimit
	}

	args = append(args, filter.Limit, filter.Offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := evr.db.Query(query, args...)

	if err != nil {
		log.Println("Error while retriving routes: ", err)
		return nil, err
	}

	for rows.Next() {
		route := er.RoutesRes{}

		err := rows.Scan(
			&route.Id,
			&route.StartPoint,
			&route.EndPoint,
			&route.Description,
		)

		if err != nil {
			log.Println("Error while scanning all routes: ", err)
			return nil, err
		}

		routes.Routes = append(routes.Routes, &route)
	}

	log.Println("Successfully fetched all routes")

	return &routes, nil
}

func (evr *EvacuationRoutesRepo) Update(req *er.RoutesUpdateReq) (*er.RoutesRes, error) {
	route := er.RoutesRes{}

	query := `
	UPDATE
		evacuation_routes
	SET 
		updated_at = now()
	`

	var conditions []string
	var args []interface{}

	if req.Route.StartPoint != "" && req.Route.StartPoint != "string" {
		conditions = append(conditions, " start_point = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Route.StartPoint)
	}
	if req.Route.EndPoint != "" && req.Route.EndPoint != "string" {
		conditions = append(conditions, " end_point = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Route.EndPoint)
	}
	if req.Route.Description != "" && req.Route.Description != "string" {
		conditions = append(conditions, " description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Route.Description)
	}
	if len(conditions) > 0 {
		query += strings.Join(conditions, ", ")
	}

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	query += `
	RETURNING
		id,
		start_point,
		end_point,
		description	
	`
	args = append(args, req.Id)

	row := evr.db.QueryRow(query, args...)

	err := row.Scan(
		&route.Id,
		&route.StartPoint,
		&route.EndPoint,
		&route.Description,
	)

	if err != nil {
		log.Println("Error while updating route: ", err)
		return nil, err
	}

	log.Println("Successfully updated route")

	return &route, nil
}

func (evr *EvacuationRoutesRepo) Delete(id *er.ById) (*er.Void, error) {
	void := er.Void{}

	query := `
	UPDATE 
		evacuation_routes
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := evr.db.Exec(query, id.Id)

	if err != nil {
		log.Println("Error while deleting route: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("route with id %s not found", id.Id)
	}

	log.Println("Successfully deleted route")

	return &void, nil
}
