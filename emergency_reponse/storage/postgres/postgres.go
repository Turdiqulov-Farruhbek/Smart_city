package postgres

import (
	"database/sql"
	"fmt"

	"emergency_response-service/config"
	"emergency_response-service/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db        *sql.DB
	AlertS    storage.EmergencyAlertsI
	DispatchS storage.ResourceDispatchesI
	DrillS    storage.EmergencyDrillsI
	FacilityS storage.EmergencyFacilitiesI
	IncidentS storage.EmergencyIncidentsI
	ResourceS storage.EmergencyResourcesI
	RouteS    storage.EvacuationRoutesI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	alert := NewEmergencyAlertsRepo(db)
	dispacht := NewResourceDispatchesRepo(db)
	drill := NewEmergencyDrillsRepo(db)
	facility := NewEmergencyFacilitiesRepo(db)
	incident := NewEmergencyIncidentsRepo(db)
	resource := NewEmergencyResourcesRepo(db)
	route := NewEvacuationRoutesRepo(db)

	return &Storage{
		Db:        db,
		AlertS:    alert,
		DispatchS: dispacht,
		DrillS:    drill,
		FacilityS: facility,
		IncidentS: incident,
		ResourceS: resource,
		RouteS:    route,
	}, nil
}
