package main

import (
	cf "emergency_response-service/config"
	er "emergency_response-service/genproto/emergency_response"
	"emergency_response-service/service"
	"emergency_response-service/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()

	db, err := postgres.NewPostgresStorage(config)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", config.EMERGENCY_RESPONSE_PORT)

	if err != nil {
		log.Fatalf("Failed to listen tcp: %v", err)
	}

	s := grpc.NewServer()

	er.RegisterAlertServiceServer(s, service.NewEmergencyAlertsService(db))
	er.RegisterEmergencyDrillServiceServer(s, service.NewEmergencyDrillsService(db))
	er.RegisterEmergencyFacilityServiceServer(s, service.NewEmergencyFacilitiesService(db))
	er.RegisterEmergencyIncidentServiceServer(s, service.NewEmergencyIncidentsService(db))
	er.RegisterEmergencyResourceServiceServer(s, service.NewEmergencyResourcesService(db))
	er.RegisterEvacuationRouteServiceServer(s, service.NewEvacuationRoutesService(db))
	er.RegisterResourceDispatcheServiceServer(s, service.NewResourceDispatchesService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}