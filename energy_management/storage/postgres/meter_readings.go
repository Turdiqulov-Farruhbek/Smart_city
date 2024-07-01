package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	pb "gitlab.com/acumen931026/energy_management/genproto/energy_managment"
)

type MeterReadingsStorage struct {
	db *sql.DB
}

func NewMeterReadingsStorage(db *sql.DB) *MeterReadingsStorage {
	return &MeterReadingsStorage{db: db}
}

func (s *MeterReadingsStorage) CreateMeterReading(req *pb.MeterReading) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
	INSERT INTO meter_readings(reading_id,meter_id,reading_value,timestamp)
	VALUES($1,$2,$3,$4)
	RETURNING reading_id, meter_id, reading_value, timestamp;
	`

	_, err := s.db.Exec(query, id, req.MeterId, req.ReadingValue, req.Time)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *MeterReadingsStorage) GetMeterReading(req *pb.ById) (*pb.GetMeterReadingWithEnergy, error) {
	query := `
        SELECT 
            m.reading_id, 
            m.meter_id, 
            m.reading_value, 
        	e.meter_id, 
            e.building_id, 
            e.meter_type, 
            e.created_at
        FROM 
            meter_readings AS m 
        JOIN 
            energy_meters AS e
        ON 
            m.meter_id = e.meter_id
        WHERE 
            m.reading_id = $1
    `

	var resp pb.GetMeterReadingWithEnergy
	resp.Energy = &pb.EnergyMeter{}
	err := s.db.QueryRow(query, req.Id).Scan(
		&resp.ReadingId,
		&resp.MeterId,
		&resp.ReadingValue,
		&resp.Time,
		&resp.Energy.MeterId,
		&resp.Energy.BuildingId,
		&resp.Energy.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *MeterReadingsStorage) DeleteMeterReading(req *pb.ById) (*pb.Void, error) {
	query := `
	UPDATE meter_readings 
		SET deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE meter_id = $1
	`
	_, err := s.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}

func (s *MeterReadingsStorage) UpdateMeterReading(req *pb.MeterReadingUpdate) (*pb.Void, error) {
	query := `
	UPDATE meter_readings
	SET meter_id = $1, reading_value = $2
	WHERE reading_id = $3
	`
	var resp pb.MeterReading
	err := s.db.QueryRow(query, req.MeterId, req.ReadingValue, req.ReadingId).
		Scan(&resp.ReadingId, &resp.MeterId, &resp.ReadingValue)
	if err != nil {
		return nil, err
	}
	return &pb.Void{}, nil
}
func (s *MeterReadingsStorage) GetHourlyEnergyForBuilding(req *pb.ByHour) (*pb.EnergyReportBuilding, error) {
    meterID := req.BuildingId
    var startTime time.Time
    var err error

    if req.Time == "" {
        startTime = time.Now()
    } else {
        startTime, err = time.Parse(time.RFC3339, req.Time)
        if err != nil {
            return nil, fmt.Errorf("invalid time format: %v", err)
        }
    }

    startOfDay := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())

    query := `
       SELECT meter_id, COALESCE(SUM(reading_value), 0) / 24 AS average_hourly_energy_value
		FROM meter_readings
		WHERE meter_id = $1
		GROUP BY meter_id;

    `

    var resp pb.EnergyReportBuilding
    err = s.db.QueryRow(query, meterID).
        Scan(&resp.BuildingId, &resp.EnergyValue)
    if err != nil {
        return nil, err
    }

    resp.Time = startOfDay.Format(time.RFC3339)
    return &resp, nil
}


func (s *MeterReadingsStorage) GetDailyEnergyForBuilding(req *pb.ById) (*pb.EnergyReportBuilding, error) {
	var resp pb.EnergyReportBuilding

	query := `
		SELECT meter_id, COALESCE(SUM(reading_value), 0) as energy_value
		FROM meter_readings
		WHERE meter_id = $1
		GROUP BY meter_id
	`

	rows, err := s.db.Query(query, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var buildingID string
		var energyValue float32

		err := rows.Scan(&buildingID, &energyValue)
		if err != nil {
			return nil, err
		}

		resp.BuildingId = buildingID
		resp.EnergyValue = energyValue
	}

	return &resp, nil
}
