package handlers

import (
	"context"
	"net/http"
	"strconv"

	emergency "gateway/genproto/emergency_response"

	"github.com/gin-gonic/gin"
)

// CreateIncidents handles POST /emergency/incidents
// @Summary Report a new incident
// @Description Report a new incident with the provided details.
// @Tags emergency
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param incident body emergency.IncidentsCreateReq true "Incident Data"
// @Success 201 {object} emergency.IncidentsRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /emergency/incidents [post]
func (h *EmergencyHandler) CreateIncidents(c *gin.Context) {
	var req emergency.IncidentsCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.EmergencyIncident.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// CreateRecourcesDispatch handles POST /emergency/resources/:resourceId/dispatch
// @Summary Dispatch a resource
// @Description Dispatch a resource to an incident or location.
// @Tags emergency
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param resourceId path string true "Resource ID"
// @Param dispatch body emergency.DispatchesCreateReq true "Resource Dispatch Data"
// @Success 201 {object} emergency.DispatchesRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /emergency/resources/{resourceId}/dispatch [post]
func (h *EmergencyHandler) CreateRecourcesDispatch(c *gin.Context) {
	var req emergency.DispatchesCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.ResourceDispatch.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// CreateAlertsAreas handles POST /emergency/alerts/areas/:areaId
// @Summary Send an alert to an area
// @Description Send an emergency alert to a specific area.
// @Tags emergency
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param areaId path string true "Area ID"
// @Param alert body emergency.AlertsCreateReq true "Alert Data"
// @Success 201 {object} emergency.AlertsRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /emergency/alerts/areas/{areaId} [post]
func (h *EmergencyHandler) CreateAlertsAreas(c *gin.Context) {
	var req emergency.AlertsCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.AlertService.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetIncidentsActive handles GET /emergency/incidents/active
// @Summary Get active incidents
// @Description Retrieve all currently active incidents.
// @Tags emergency
// @Produce json
// @Security BearerAuth
// @Success 200 {array} emergency.IncidentsGetByIdRes
// @Failure 500 {object} string "Internal Server Error"
// @Router /emergency/incidents/active [get]
func (h *EmergencyHandler) GetIncidentsActive(c *gin.Context) {
	resp, err := h.EmergencyIncident.GetById(context.Background(), &emergency.ById{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetIncidentsGetAll handles GET /emergency/resources/available/:emergencyTypeId
// @Summary Get available resources by emergency type
// @Description Get all available resources for a specific type of emergency.
// @Tags emergency
// @Produce json
// @Security BearerAuth
// @Param emergencyTypeId path string true "Emergency Type ID"
// @Success 200 {array} emergency.ResourcesGetAllRes
// @Failure 500 {object} string "Internal Server Error"
// @Router /emergency/resources/available/{emergencyTypeId} [get]
func (h *EmergencyHandler) GetEmergencyAll(c *gin.Context) {
	limit := c.Param("limit")
	offset := c.Param("offset")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	req := &emergency.Filter{Limit: int32(limitInt), Offset: int32(offsetInt)}

	resp, err := h.EmergencyResource.GetAll(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetIncidentsGetAll handles GET /emergency/resources/available/:emergencyTypeId
// @Summary Get available resources by emergency type
// @Description Get all available resources for a specific type of emergency.
// @Tags emergency
// @Produce json
// @Security BearerAuth
// @Param emergencyTypeId path string true "Emergency Type ID"
// @Success 200 {array} emergency.ResourcesGetByIdRes
// @Failure 500 {object} string "Internal Server Error"
// @Router /emergency/resources/available/{emergencyTypeId} [get]
func (h *EmergencyHandler) GetIncidentsGetAll(c *gin.Context) {
	id := c.Param("emergencyTypeId")
	req := &emergency.ById{Id: id}

	resp, err := h.EmergencyResource.GetById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetEvacuationRoutes handles GET /emergency/evacuation/routes/:areaId
// @Summary Get evacuation routes for an area
// @Description Get recommended evacuation routes for a specific area.
// @Tags emergency
// @Produce json
// @Security BearerAuth
// @Param areaId path string true "Area ID"
// @Success 200 {array} emergency.RoutesGetByIdRes
// @Failure 500 {object} string "Internal Server Error"
// @Router /emergency/evacuation/routes/{areaId} [get]
func (h *EmergencyHandler) GetEvacuationRoutes(c *gin.Context) {
	id := c.Param("areaId")
	req := &emergency.ById{Id: id}

	resp, err := h.EvacuationRoute.GetById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateIncidentsStatus handles PUT /emergency/incidents/:incidentId/status
// @Summary Update incident status
// @Description Update the status of an incident by its ID.
// @Tags emergency
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param incidentId path string true "Incident ID"
// @Param status body emergency.IncidentsUpdateReq true "Incident Status Update"
// @Success 200 {object} emergency.IncidentsRes
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /emergency/incidents/{incidentId}/status [put]
func (h *EmergencyHandler) UpdateIncidentsStatus(c *gin.Context) {
	id := c.Param("incidentId")
	var req emergency.IncidentsUpdateReq
	req.Id = id

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.EmergencyIncident.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
