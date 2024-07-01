package handlers

import (
	"context"
	"net/http"

	environment "gateway/genproto/enviromental_monitoring"

	"github.com/gin-gonic/gin"
)


// CreatePlantsSpeciesIdRegister handles POST /environment/plants/:speciesId/register
// @Summary Register a plant species
// @Description Register a new plant species.
// @Tags environment
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param speciesId path string true "Species ID"
// @Param plant body environment.PlantRegistry true "Plant Species Registration Data"
// @Success 201 {object} environment.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /environment/plants/{speciesId}/register [post]
func (h *EnvironmentalMonitoringHandler) CreatePlantsPeciesIdRegister(c *gin.Context) {
	var req environment.PlantRegistry
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.PlantRegistry.RegisterPlants(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetAirQualityStationsStationIdCurrent handles GET /environment/air-quality/stations/:stationId/current
// @Summary Get current air quality
// @Description Retrieve current air quality data from a specific station.
// @Tags environment
// @Produce json
// @Security BearerAuth
// @Param stationId path string true "Station ID"
// @Success 200 {object} environment.Station
// @Failure 500 {object} string "Internal Server Error"
// @Router /environment/air-quality/stations/{stationId}/current [get]
func (h *EnvironmentalMonitoringHandler) GetAirGualityStationsStationIdCurrent(c *gin.Context) {
	id := c.Param("stationId")
	req := &environment.ById{Id: id}

	resp, err := h.AirQualityStation.GetStation(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAirQualityCityForecast handles GET /environment/air-quality/city/forecast
// @Summary Get city air quality forecast
// @Description Retrieve air quality forecast data for the city.
// @Tags environment
// @Produce json
// @Security BearerAuth
// @Success 200 {object} environment.AirQualityList
// @Failure 500 {object} string "Internal Server Error"
// @Router /environment/air-quality/city/forecast [get]
func (h *EnvironmentalMonitoringHandler) GetAirQualityCityForecast(c *gin.Context) {
	resp, err := h.AirQualityReading.GetCityWideAirQuality(context.Background(), &environment.Void{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetNoiseZonesZoneIdLevels handles GET /environment/noise/zones/:zoneId/levels
// @Summary Get noise levels in a zone
// @Description Retrieve noise levels for a specific zone.
// @Tags environment
// @Produce json
// @Security BearerAuth
// @Param zoneId path string true "Zone ID"
// @Success 200 {object} environment.NoiseArea
// @Failure 500 {object} string "Internal Server Error"
// @Router /environment/noise/zones/{zoneId}/levels [get]
func (h *EnvironmentalMonitoringHandler) GetNoiseZonesZoneIdLevels(c *gin.Context) {
	id := c.Param("zoneId")
	req := &environment.ById{Id: id}

	resp, err := h.NoiseMonitoringAre.GetNoiseMonitoringArea(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRecyclingCentersAreaId handles GET /environment/recycling/centers/:areaId
// @Summary Get recycling centers by area
// @Description Retrieve information about recycling centers in a specific area.
// @Tags environment
// @Produce json
// @Security BearerAuth
// @Param areaId path string true "Area ID"
// @Success 200 {object} environment.RecyclingCenterList
// @Failure 500 {object} string "Internal Server Error"
// @Router /environment/recycling/centers/{areaId} [get]
func (h *EnvironmentalMonitoringHandler) GetRecyclingCentersAreaId(c *gin.Context) {
	id := c.Param("areaId")
	req := &environment.RecyclingCenterFilter{CenterId: id}

	resp, err := h.RecyclingCenter.GetRecyclingCenters(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetGreenSpacesSpaceIdStatus handles GET /environment/green-spaces/:spaceId/status
// @Summary Get green space status
// @Description Retrieve status information for a specific green space.
// @Tags environment
// @Produce json
// @Security BearerAuth
// @Param spaceId path string true "Space ID"
// @Success 200 {object} environment.GreenSpaceList
// @Failure 500 {object} string "Internal Server Error"
// @Router /environment/green-spaces/{spaceId}/status [get]
func (h *EnvironmentalMonitoringHandler) GetGreenSpacesSpaceIdStatus(c *gin.Context) {
	id := c.Param("spaceId")
	req := &environment.GreenSpaceFilter{SpaceId: id}

	resp, err := h.GreenSpace.GetGreenSpaces(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
