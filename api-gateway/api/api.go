package api

import (
	handler "gateway/api/handlers"
	// "gateway/api/middlerware"
	
	_ "gateway/docs"
	_ "gateway/genproto/auth"
	_ "gateway/genproto/citizen"
	_ "gateway/genproto/city_planning"
	_ "gateway/genproto/emergency_response"
	_ "gateway/genproto/energy_management"
	_ "gateway/genproto/enviromental_monitoring"
	_ "gateway/genproto/transport"

	// "log"
	// "github.com/casbin/casbin/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

// @title Smart City Management System API
// @version 1.0
// @description API for Smart City Management System resources
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(/*AuthConn, citizenConn, cityPlaningConn, emergencyConn, energyConn, enviromentConn, transportconn*/citizenConn *grpc.ClientConn) *gin.Engine {
	// authH := handlers.NewAuthHandler(conn)
	citizenH := handler.NewCitizenHandler(citizenConn)
	// cityPlanningH := handler.NewCityPlanningHandler(cityPlaningConn)
	// emergencyH := handler.NewEmergencyHandler(emergencyConn)
	// energyH := handler.NewEnergyHandler(energyConn)
	// enviromentalH := handler.NewEnvironmentalMonitoringHandler(enviromentConn)
	// transportH := handler.NewTransportationHandler(transportconn)

	router := gin.Default()
	
	// enforcer, err := casbin.NewEnforcer("gateway/casbin/model.conf","gateway/casbin/policy.csv")
	// if err!= nil {
    //     log.Fatal(err)
    // }
	// router.Use(middlerware.NewAuth(enforcer))


	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Adjust for your specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	citizen := router.Group("/citizens")
	{
		citizen.POST("/create", citizenH.CreateCitizen)
		citizen.POST("/:citizenId/requests", citizenH.CreateCitizensRequests)
		citizen.POST("/:citizenId/documents", citizenH.CreateCitizensDocuments)

		citizen.GET("/:citizenId", citizenH.GetCitizen)
		citizen.GET("/:citizenId/requests", citizenH.GetCitizensrequests)

		citizen.PUT("/:citizenId", citizenH.UpdateCitizen)

	}

	// cityPlanning := router.Group("/planning")
	// {
	// 	cityPlanning.POST("/proposals", cityPlanningH.CreatePlanningProposals)
	// 	cityPlanning.POST("/proposals/:proposalId/feedback", cityPlanningH.CreatePlanningProposalsFeedback)
	// 	cityPlanning.POST("/land-use/:parcelId/change-request", cityPlanningH.CreatePlanningProposalsChangeRequest)

	// 	cityPlanning.GET("/zones/:zoneId/current", cityPlanningH.GetPlanningProposalsCurrentZones)
	// 	cityPlanning.GET("/demographics/:areaId", cityPlanningH.GetPlanningProposalsDemographicsAria)
	// 	cityPlanning.GET("/infrastructure/:infrastructureId/status", cityPlanningH.GetPlanningProposalsInfrastructure)
	// 	cityPlanning.GET("/permits/:permitTypeId/requirements", cityPlanningH.GetPlanningProposalsPermitsRequirements)
	// }

	// emergency := router.GROUP("/emergency")
	// {
	// 	emergency.POST("/incidents", emergencyH.CreateIncidents)
	// 	emergency.POST("/resources/:resourceId/dispatch", emergencyH.CreateRecourcesDispatch)
	// 	emergency.POST("/alerts/areas/:areaId", emergencyH.CreateAlertsAreas)

	// 	emergency.GET("/incidents/active", emergencyH.GetIncidentsActive)
	// 	emergency.GET("/resources/available/:emergencyTypeId", emergencyH.GetIncidentsGetAll)
	// 	emergency.GET("/", emergencyH.GetEmergencyAll)
	// 	emergency.GET("/evacuation/routes/:areaId", emergencyH.GetEvacuationRoutes)

	// 	emergency.PUT("/incidents/:incidentId/status", emergencyH.UpdateIncidentsStatus)
	// }

	// energy := router.GROUP("/energy")
	// {
	// 	energy.POST("/meters/:meterId/readings", energyH.CreateEnergyMeterReadings)
	// 	energy.GET("/consumption/city/daily", energyH.GetConsumptionCityDaily)
	// 	energy.GET("/consumption/buildings/:buildingId/hourly", energyH.GetConsumptionBuildingIdHourly)

	// }

	// enviroment := router.GET("/enviroment")
	// {
	// 	enviroment.POST("/environment/plants/:speciesId/register", enviromentalH.CreatePlantsPeciesIdRegister)

	// 	enviroment.GET("/air-quality/stations/:stationId/current", enviromentalH.GetAirGualityStationsStationIdCurrent)
	// 	enviroment.GET("/air-quality/city/forecast", enviromentalH.GetAirQualityCityForecast)
	// 	enviroment.GET("/noise/zones/:zoneId/levels", enviromentalH.GetNoiseZonesZoneIdLevels)
	// 	enviroment.GET("/recycling/centers/:areaId", enviromentalH.GetRecyclingCentersAreaId)
	// 	enviroment.GET("/green-spaces/:spaceId/status", enviromentalH.GetGreenSpacesSpaceIdStatus)

	// }

	// transport := router.GROUP("/transport")
	// {
	// 	transport.POST("/parking-lots/:lotId/reserve", transportH.CreateParkingLotsLotIdReserve)
	// 	transport.POST("/incidents/report", transportH.CreateIncidentsReport)
	// 	transport.POST("/roads/:roadId/signals", transportH.CreateRoadsRoadIdSignals)

	// 	transport.GET("/routes", transportH.GetRoutes)
	// 	transport.GET("/routes/:routeId/schedule", transportH.GetRoutesRouteIdSchedule)
	// 	transport.GET("/parking-lots", transportH.GetParkingLots)
	// 	transport.GET("/parking-lots/:lotId/status", transportH.GetParkingLotsLotIdStatus)
	// 	transport.GET("/roads/:roadId/traffic", transportH.GetRoadsRoadIdTraffic)
	// 	transport.GET("/maintenance/schedule", transportH.GetMaintenanceSchedule)

	// }

	return router
}
