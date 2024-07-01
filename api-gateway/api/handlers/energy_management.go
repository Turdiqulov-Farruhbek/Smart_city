package handlers

import (
	"context"
	"net/http"

	energy "gateway/genproto/energy_management"

	"github.com/gin-gonic/gin"
)


// CreateEnergyMeterReadings handles POST /energy/meters/:meterId/readings
// @Summary Record meter readings
// @Description Record energy meter readings for a specific meter.
// @Tags energy
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param meterId path string true "Meter ID"
// @Param reading body energy.MeterReading true "Meter Reading Data"
// @Success 201 {object} energy.Void
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /energy/meters/{meterId}/readings [post]
func (h *EnergyHandler) CreateEnergyMeterReadings(c *gin.Context) {
	var req energy.MeterReading
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	resp, err := h.MeterReading.CreateMeterReading(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}


// GetConsumptionCityDaily handles GET /energy/consumption/city/daily
// @Summary Get daily city energy consumption
// @Description Retrieve daily energy consumption data for the city.
// @Tags energy
// @Produce json
// @Security BearerAuth
// @Success 200 {object} energy.EnergyMeter
// @Failure 500 {object} string "Internal Server Error"
// @Router /energy/consumption/city/daily [get]
func (h *EnergyHandler) GetConsumptionCityDaily(c *gin.Context) {
	resp, err := h.EnergyConsumption.GetEnergyMeter(context.Background(), &energy.ById{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetConsumptionBuildingIdHourly handles GET /energy/consumption/buildings/:buildingId/hourly
// @Summary Get hourly building energy consumption
// @Description Retrieve hourly energy consumption data for a specific building.
// @Tags energy
// @Produce json
// @Security BearerAuth
// @Param buildingId path string true "Building ID"
// @Success 200 {object} energy.Building
// @Failure 500 {object} string "Internal Server Error"
// @Router /energy/consumption/buildings/{buildingId}/hourly [get]
func (h *EnergyHandler) GetConsumptionBuildingIdHourly(c *gin.Context) {
	id := c.Param("buildingId")
	req := &energy.ById{Id: id}

	resp, err := h.BuildingService.GetBuilding(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}


