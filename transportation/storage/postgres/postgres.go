package postgres

import (
	"database/sql"
	"fmt"
	"gitlab.com/acumen/smart_city/transportation/storage"

	"golang.org/x/exp/slog"

	_ "github.com/lib/pq"
	"gitlab.com/acumen/smart_city/transportation/config"
	"gitlab.com/acumen/smart_city/transportation/config/logger"
)

type Storage struct {
	Db                *sql.DB
	Logger            *logger.Logger
	IncidentS         storage.IncidentI
	ParkingLotS       storage.ParkingLotI
	ParkingSpaceS     storage.ParkingSpaceI
	RoadS             storage.RoadI
	RouteScheduleS    storage.RouteScheduleI
	RouteS            storage.RouteI
	TrafficConditionS storage.TrafficConditionI
	VehicleS          storage.VehicleI
}

func NewPostgresStorage(config config.Config, logger *logger.Logger) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST,
		config.DB_USER,
		config.DB_NAME,
		config.DB_PASSWORD,
		config.DB_PORT,
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	slog.Info("connected to db")

	return &Storage{
		Db:                db,
		IncidentS:         NewIncidentsRepo(db),
		ParkingLotS:       NewParkingLotsRepo(db),
		ParkingSpaceS:     NewParkingSpacesRepo(db),
		RoadS:             NewRoadsRepo(db),
		RouteScheduleS:    NewRouteScheduleRepo(db),
		RouteS:            NewRoadsRepo(db),
		TrafficConditionS: NewTrafficConditionsRepo(db),
		VehicleS:          NewVehiclesRepo(db),
	}, nil
}
