package postgres

import (
	"database/sql"

	"github.com/google/uuid"
	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
)

type EnergyMetersStorage struct {
	db *sql.DB
}

func NewEnergyMetersStorage(db *sql.DB) *EnergyMetersStorage {
	return &EnergyMetersStorage{db: db}
}

func (s *EnergyMetersStorage) CreateEnergyMeter(req *pb.EnergyMeterCreate) (*pb.EnergyMeter, error) {
	id := uuid.NewString()

	query := `
    INSERT INTO energy_meters (meter_id, building_id, meter_type)
    VALUES ($1, $2, $3)
    RETURNING meter_id, building_id, meter_type;
    `

	var resp pb.EnergyMeter
	err := s.db.QueryRow(query, id, req.BuildingId, req.MeterType).
		Scan(&resp.MeterId, &resp.BuildingId, &resp.MeterType)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}


func (s *EnergyMetersStorage) UpdateEnergyMeter(req *pb.EnergyMeterUpdate) (*pb.EnergyMeter, error) {
	query := `
	UPDATE energy_meters
	SET building_id = $1, meter_type = $2,updated_at = now()
	WHERE meter_id = $3
	RETURNING meter_id, building_id, meter_type;
	`
	var resp pb.EnergyMeter
	err := s.db.QueryRow(query, req.BuildingId, req.MeterType, req.MeterId).
		Scan(&resp.MeterId, &resp.BuildingId, &resp.MeterType)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *EnergyMetersStorage) GetEnergyMeter(req *pb.ById) (*pb.GetEnergyWithBuildings, error) {
	query := `
        SELECT 
            e.meter_id, 
            e.building_id, 
            e.meter_type, 
            e.created_at, 
            b.building_id, 
            b.address, 
            b.building_type, 
            b.total_area, 
            b.created_at
        FROM 
            energy_meters AS e 
        JOIN 
            buildings AS b 
        ON 
            e.building_id = b.building_id
        WHERE 
            e.meter_id = $1
    `

	var resp pb.GetEnergyWithBuildings
	resp.Meter = &pb.EnergyMeter{}
	resp.Building = &pb.Building{}

	err := s.db.QueryRow(query, req.Id).Scan(
		&resp.Meter.MeterId,
		&resp.Meter.BuildingId,
		&resp.Meter.MeterType,
		&resp.Meter.CreatedAt,
		&resp.Building.BuildingId,
		&resp.Building.Address,
		&resp.Building.BuildingType,
		&resp.Building.TotalArea,
		&resp.Building.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *EnergyMetersStorage) DeleteEnergyMeter(req *pb.ById) (*pb.Void, error) {
	query := `
	UPDATE energy_meters 
		SET deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE meter_id = $1
	`
	_, err := s.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
