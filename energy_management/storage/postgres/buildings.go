package postgres

import (
	"database/sql"

	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"

	"github.com/google/uuid"
)

type BuildingsStorage struct {
	db *sql.DB
}

func NewBuildingsStorage(db *sql.DB) *BuildingsStorage {
	return &BuildingsStorage{db: db}
}

func (s *BuildingsStorage) CreateBuilding(req *pb.BuildingCreate) (*pb.Building, error) {
	id := uuid.NewString()
	query := `
        INSERT INTO buildings (building_id, address, building_type, total_area)
        VALUES ($1, $2, $3, $4)
        RETURNING building_id, address, building_type, total_area
    `
	var resp pb.Building
	err := s.db.QueryRow(query, id, req.Address, req.BuildingType, req.TotalArea).
		Scan(&resp.BuildingId, &resp.Address, &resp.BuildingType, &resp.TotalArea)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *BuildingsStorage) UpdateBuilding(req *pb.BuildingUpdate) (*pb.Building, error) {
	query := `
	UPDATE buildings
	SET address = $1, building_type = $2, total_area = $3, updated_at = now()
	WHERE building_id = $4
	RETURNING building_id, address, building_type, total_area
	`
	var resp pb.Building
	err := s.db.QueryRow(query, req.Address, req.BuildingType, req.TotalArea, req.Id).
		Scan(&resp.BuildingId, &resp.Address, &resp.BuildingType, &resp.TotalArea)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BuildingsStorage) GetBuilding(req *pb.ById) (*pb.Building, error) {
	query := `
	SELECT building_id, address, building_type, total_area
	FROM buildings
	WHERE building_id = $1
	`
	var resp pb.Building
	err := s.db.QueryRow(query, req.Id).Scan(&resp.BuildingId, &resp.Address, &resp.BuildingType, &resp.TotalArea)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *BuildingsStorage) DeleteBuilding(req *pb.ById) (*pb.Void, error) {
	query := `
	UPDATE buildings 
		SET deleted_at = EXTRACT(EPOCH FROM NOW()) 
	WHERE building_id = $1
	`
	_, err := s.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
