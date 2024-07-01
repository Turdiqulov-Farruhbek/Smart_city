package handlers

import (
	"gateway/genproto/auth"
	"gateway/genproto/citizen"
	"gateway/genproto/city_planning"
	"gateway/genproto/emergency_response"
	"gateway/genproto/energy_management"
	"gateway/genproto/enviromental_monitoring"
	"gateway/genproto/transport"

	"google.golang.org/grpc"
)

type AuthHandler struct {
	Auth auth.AuthServiceClient
}

func NewAuthHandler(authConn *grpc.ClientConn) *AuthHandler {
	return &AuthHandler{
		Auth: auth.NewAuthServiceClient(authConn),
	}
}

type CitizenHandler struct {
	Citizen             citizen.CitizenServiceClient
	CitizenDocument     citizen.CitizenDocumentServiceClient
	CitizenNotification citizen.CitizenNotificationServiceClient
	CitizenRequest      citizen.CitizenRequestServiceClient
}

func NewCitizenHandler(citizenConn *grpc.ClientConn) *CitizenHandler {
	return &CitizenHandler{
		Citizen:             citizen.NewCitizenServiceClient(citizenConn),
		CitizenDocument:     citizen.NewCitizenDocumentServiceClient(citizenConn),
		CitizenNotification: citizen.NewCitizenNotificationServiceClient(citizenConn),
		CitizenRequest:      citizen.NewCitizenRequestServiceClient(citizenConn),
	}
}

type CityPlanningHandler struct {
	CityPlanning        city_planning.CityZoneServiceClient
	BuildingPermit      city_planning.BuildingPermitServiceClient
	DemographicData     city_planning.DemographicDataServiceClient
	DevelopmentScenario city_planning.DevelopmentScenarioServiceClient
	InfrastuctureAsset  city_planning.InfrastuctureAssetServiceClient
	PlanningProposal    city_planning.PlanningProposalServiceClient
	ProposalFeedback    city_planning.ProposalFeedbackServiceClient
}

func NewCityPlanningHandler(cityPlanningConn *grpc.ClientConn) *CityPlanningHandler {
	return &CityPlanningHandler{
		CityPlanning:        city_planning.NewCityZoneServiceClient(cityPlanningConn),
		BuildingPermit:      city_planning.NewBuildingPermitServiceClient(cityPlanningConn),
		DemographicData:     city_planning.NewDemographicDataServiceClient(cityPlanningConn),
		DevelopmentScenario: city_planning.NewDevelopmentScenarioServiceClient(cityPlanningConn),
		InfrastuctureAsset:  city_planning.NewInfrastuctureAssetServiceClient(cityPlanningConn),
		PlanningProposal:    city_planning.NewPlanningProposalServiceClient(cityPlanningConn),
		ProposalFeedback:    city_planning.NewProposalFeedbackServiceClient(cityPlanningConn),
	}
}

type EmergencyHandler struct {
	AlertService      emergency_response.AlertServiceClient
	ResourceDispatch  emergency_response.ResourceDispatchServiceClient
	EmergencyDrill    emergency_response.EmergencyDrillServiceClient
	EmergencyFacility emergency_response.EmergencyFacilityServiceClient
	EmergencyIncident emergency_response.EmergencyIncidentServiceClient
	EmergencyResource emergency_response.EmergencyResourceServiceClient
	EvacuationRoute   emergency_response.EvacuationRouteServiceClient
}

func NewEmergencyHandler(emergencyConn *grpc.ClientConn) *EmergencyHandler {
	return &EmergencyHandler{
		AlertService:      emergency_response.NewAlertServiceClient(emergencyConn),
		ResourceDispatch:  emergency_response.NewResourceDispatchServiceClient(emergencyConn),
		EmergencyDrill:    emergency_response.NewEmergencyDrillServiceClient(emergencyConn),
		EmergencyFacility: emergency_response.NewEmergencyFacilityServiceClient(emergencyConn),
		EmergencyIncident: emergency_response.NewEmergencyIncidentServiceClient(emergencyConn),
		EmergencyResource: emergency_response.NewEmergencyResourceServiceClient(emergencyConn),
		EvacuationRoute:   emergency_response.NewEvacuationRouteServiceClient(emergencyConn),
	}
}

type EnergyHandler struct {
	BuildingService   energy_management.BuildingServiceClient
	EnergyConsumption energy_management.EnenrgyMeterServiceClient
	MeterReading      energy_management.MeterReadingServiceClient
}

func NewEnergyHandler(energyConn *grpc.ClientConn) *EnergyHandler {
	return &EnergyHandler{
		BuildingService:   energy_management.NewBuildingServiceClient(energyConn),
		EnergyConsumption: energy_management.NewEnenrgyMeterServiceClient(energyConn),
		MeterReading:      energy_management.NewMeterReadingServiceClient(energyConn),
	}
}

type EnvironmentalMonitoringHandler struct {
	AirQualityReading  enviromental_monitoring.AirQualityReadingServiceClient
	AirQualityStation  enviromental_monitoring.AirQualityStationServiceClient
	GreenSpace         enviromental_monitoring.GreenSpaceServiceClient
	NoiseLevelReading  enviromental_monitoring.NoiseLevelReadingServiceClient
	NoiseMonitoringAre enviromental_monitoring.NoiseMonitoringAreServiceClient
	PlantRegistry      enviromental_monitoring.PlantRegistryServiceClient
	RecyclingCenter    enviromental_monitoring.RecyclingCenterServiceClient
	WasteCollection    enviromental_monitoring.WasteCollectionServiceClient
}

func NewEnvironmentalMonitoringHandler(enviromentalMonitoringConn *grpc.ClientConn) *EnvironmentalMonitoringHandler {
	return &EnvironmentalMonitoringHandler{
		AirQualityReading:  enviromental_monitoring.NewAirQualityReadingServiceClient(enviromentalMonitoringConn),
		AirQualityStation:  enviromental_monitoring.NewAirQualityStationServiceClient(enviromentalMonitoringConn),
		GreenSpace:         enviromental_monitoring.NewGreenSpaceServiceClient(enviromentalMonitoringConn),
		NoiseLevelReading:  enviromental_monitoring.NewNoiseLevelReadingServiceClient(enviromentalMonitoringConn),
		NoiseMonitoringAre: enviromental_monitoring.NewNoiseMonitoringAreServiceClient(enviromentalMonitoringConn),
		PlantRegistry:      enviromental_monitoring.NewPlantRegistryServiceClient(enviromentalMonitoringConn),
		RecyclingCenter:    enviromental_monitoring.NewRecyclingCenterServiceClient(enviromentalMonitoringConn),
		WasteCollection:    enviromental_monitoring.NewWasteCollectionServiceClient(enviromentalMonitoringConn),
	}
}

type TransportHandler struct {
	Incidemt       transport.IncidemtServiceClient
	Maintance      transport.MaintanceServiceClient
	Parking        transport.ParkingServiceClient
	RoadService    transport.RoadServiceClient
	RouteService   transport.RouteServiceClient
	TrafficService transport.TrafficServiceClient
	Transport      transport.TransportSeviceClient
}

func NewTransportationHandler(transportationConn *grpc.ClientConn) *TransportHandler {
	return &TransportHandler{
		Incidemt:       transport.NewIncidemtServiceClient(transportationConn),
		Maintance:      transport.NewMaintanceServiceClient(transportationConn),
		Parking:        transport.NewParkingServiceClient(transportationConn),
		RoadService:    transport.NewRoadServiceClient(transportationConn),
		RouteService:   transport.NewRouteServiceClient(transportationConn),
		TrafficService: transport.NewTrafficServiceClient(transportationConn),
		Transport:      transport.NewTransportSeviceClient(transportationConn),
	}
}
