package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gitlab.com/acumen931026/energy_management/config"
	"gitlab.com/acumen931026/energy_management/storage"
)

type Storage struct {
	db        *sql.DB
	BuildingS storage.BuildingI
	EnergyS   storage.EnergyMeterI
	MeterS    storage.MeterI
}

func ConnectDB(config.Config) (*Storage, error) {
	cfg := config.Load()
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	buildingS := NewBuildingsStorage(db)
	energyS := NewEnergyMetersStorage(db)
	meterS := NewMeterReadingsStorage(db)
	return &Storage{
		db:        db,
		BuildingS: buildingS,
		EnergyS:   energyS,
		MeterS:    meterS,
	}, nil
}
func (s *Storage) Building() storage.BuildingI {
	if s.BuildingS == nil {
		s.BuildingS = NewBuildingsStorage(s.db)
	}
	return s.BuildingS
}

func (s *Storage) EnergyMeter() storage.EnergyMeterI {
	if s.EnergyS == nil {
		s.EnergyS = NewEnergyMetersStorage(s.db)
	}
	return s.EnergyS
}

func (s *Storage) Meter() storage.MeterI {
	if s.MeterS == nil {
		s.MeterS = NewMeterReadingsStorage(s.db)
	}
	return s.MeterS
}
