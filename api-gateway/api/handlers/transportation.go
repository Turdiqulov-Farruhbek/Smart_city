package handlers

import (
	"context"
	"net/http"

	"gateway/genproto/transport"

	"github.com/gin-gonic/gin"
)


// CreateParkingLotsLotIdReserve handles POST /transport/parking-lots/:lotId/reserve
// @Summary Reserve a parking lot
// @Description Reserve a parking lot by its ID.
// @Tags transport
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param lotId path string true "Lot ID"
// @Param reservation body transport.ParkingLotCreate true "Parking Lot Reservation Data"
// @Success 201 {object} transport.ParkingLot
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /transport/parking-lots/{lotId}/reserve [post]
func (h *TransportHandler) CreateParkingLotsLotIdReserve(c *gin.Context) {
	var req transport.ParkingLotCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Parking.CreateParkingLot(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// CreateIncidentsReport handles POST /transport/incidents/report
// @Summary Report an incident
// @Description Report a transportation-related incident.
// @Tags transport
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param incident body transport.IncidentCreate true "Incident Report Data"
// @Success 201 {object} transport.Incident
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /transport/incidents/report [post]
func (h *TransportHandler) CreateIncidentsReport(c *gin.Context) {
	var req transport.IncidentCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Incidemt.CreateIncident(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// CreateRoadsRoadIdSignals handles POST /transport/roads/:roadId/signals
// @Summary Control road signals
// @Description Create or update signals for a specific road.
// @Tags transport
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param roadId path string true "Road ID"
// @Param signals body transport.RoadCreate true "Road Signal Control Data"
// @Success 201 {object} transport.Road
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /transport/roads/{roadId}/signals [post]
func (h *TransportHandler) CreateRoadsRoadIdSignals(c *gin.Context) {
	var req transport.RoadCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.RoadService.CreateRoad(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetRoutes handles GET /transport/routes
// @Summary Get all transport routes
// @Description Retrieve all transport routes.
// @Tags transport
// @Produce json
// @Security BearerAuth
// @Success 200 {object} transport.RouteList
// @Failure 500 {object} string "Internal Server Error"
// @Router /transport/routes [get]
func (h *TransportHandler) GetRoutes(c *gin.Context) {
	resp, err := h.RouteService.GetAllTrasnportRoutes(context.Background(), &transport.RouteFilter{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRoutesRouteIdSchedule handles GET /transport/routes/:routeId/schedule
// @Summary Get route schedule
// @Description Retrieve the schedule for a specific transport route.
// @Tags transport
// @Produce json
// @Security BearerAuth
// @Param routeId path string true "Route ID"
// @Success 200 {object} transport.RouteScheduleList
// @Failure 500 {object} string "Internal Server Error"
// @Router /transport/routes/{routeId}/schedule [get]
func (h *TransportHandler) GetRoutesRouteIdSchedule(c *gin.Context) {
	id := c.Param("routeId")
	req := &transport.ById{Id: id}

	resp, err := h.RouteService.GetScheduleForRoute(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}



// GetParkingLots handles GET /transport/parking-lots
// @Summary Get all parking lots
// @Description Retrieve information about all parking lots.
// @Tags transport
// @Produce json
// @Security BearerAuth
// @Success 200 {object} transport.PakingLotList
// @Failure 500 {object} string "Internal Server Error"
// @Router /transport/parking-lots [get]
func (h *TransportHandler) GetParkingLots(c *gin.Context) {
	resp, err := h.Parking.GetAllParkingLots(context.Background(), &transport.ParkingLotFilter{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetParkingLotsLotIdStatus handles GET /transport/parking-lots/:lotId/status
// @Summary Get parking lot status
// @Description Retrieve status information for a specific parking lot.
// @Tags transport
// @Produce json
// @Security BearerAuth
// @Param lotId path string true "Lot ID"
// @Success 200 {object} transport.ParkingLot
// @Failure 500 {object} string "Internal Server Error"
// @Router /transport/parking-lots/{lotId}/status [get]
func (h *TransportHandler) GetParkingLotsLotIdStatus(c *gin.Context) {
	id := c.Param("lotId")
	req := &transport.ById{Id: id}

	resp, err := h.Parking.GetParkingStatus(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRoadsRoadIdTraffic handles GET /transport/roads/:roadId/traffic
// @Summary Get road traffic status
// @Description Retrieve traffic information for a specific road.
// @Tags transport
// @Produce json
// @Security BearerAuth
// @Param roadId path string true "Road ID"
// @Success 200 {object} transport.Road
// @Failure 500 {object} string "Internal Server Error"
// @Router /transport/roads/{roadId}/traffic [get]
func (h *TransportHandler) GetRoadsRoadIdTraffic(c *gin.Context) {
	id := c.Param("roadId")
	req := &transport.ById{Id: id}

	resp, err := h.RoadService.GetRoad(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetMaintenanceSchedule handles GET /transport/maintenance/schedule
// @Summary Get maintenance schedule
// @Description Retrieve the maintenance schedule for transport infrastructure.
// @Tags transport
// @Produce json
// @Security BearerAuth
// @Success 200 {object} transport.MaintanceScheduleList
// @Failure 500 {object} string "Internal Server Error"
// @Router /transport/maintenance/schedule [get]
func (h *TransportHandler) GetMaintenanceSchedule(c *gin.Context) {
	resp, err := h.Maintance.GetAllMaintanceSchedules(context.Background(), &transport.MaintanceScheduleFilter{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
