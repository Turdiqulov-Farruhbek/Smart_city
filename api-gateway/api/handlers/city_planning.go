package handlers


import (
	"context"
	"gateway/genproto/city_planning"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePlanningProposals handles POST /planning/proposals
// @Summary Create a new planning proposal
// @Description Create a new planning proposal with the provided details.
// @Tags city_planning
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param proposal body city_planning.PlanningProposalCreate true "Planning Proposal Data"
// @Success 201 {object} city_planning.Void
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /planning/proposals [post]
func (h *CityPlanningHandler) CreatePlanningProposals(c *gin.Context) {
	var req city_planning.PlanningProposalCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.PlanningProposal.CreatePlanningProposal(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// CreatePlanningProposalsFeedback handles POST /planning/proposals/:proposalId/feedback
// @Summary Submit feedback for a planning proposal
// @Description Submit feedback for a specific planning proposal.
// @Tags city_planning
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param proposalId path string true "Proposal ID"
// @Param feedback body city_planning.ProposalFeedbackCreate true "Proposal Feedback Data"
// @Success 201 {object} city_planning.Void
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /planning/proposals/{proposalId}/feedback [post]
func (h *CityPlanningHandler) CreatePlanningProposalsFeedback(c *gin.Context) {
	var req city_planning.ProposalFeedbackCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.ProposalFeedback.CreateProposalFeedback(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// CreatePlanningProposalsChangeRequest handles POST /planning/land-use/:parcelId/change-request
// @Summary Create a land-use change request
// @Description Create a land-use change request for a specific parcel.
// @Tags city_planning
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param parcelId path string true "Parcel ID"
// @Success 201 {object} city_planning.Void
// @Failure 400 {object} string "error": "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /planning/land-use/{parcelId}/change-request [post]
func (h *CityPlanningHandler) CreatePlanningProposalsChangeRequest(c *gin.Context) {
	var req city_planning.PlanningProposalCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.PlanningProposal.CreatePlanningProposal(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetPlanningProposalsCurrentZones handles GET /planning/zones/:zoneId/current
// @Summary Get current zone details
// @Description Get current zone details by zone ID.
// @Tags city_planning
// @Produce json
// @Security BearerAuth
// @Param zoneId path string true "Zone ID"
// @Success 200 {object} city_planning.CityZoneList
// @Failure 500 {object} string "Internal Server Error"
// @Router /planning/zones/{zoneId}/current [get]
func (h *CityPlanningHandler) GetPlanningProposalsCurrentZones(c *gin.Context) {
	id := c.Param("zoneId")
	req := &city_planning.CityZoneFilter{ZoneId: id}

	resp, err := h.CityPlanning.GetCityZone(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetPlanningProposalsDemographicsAria handles GET /planning/demographics/:areaId
// @Summary Get demographics details
// @Description Get demographics details for a specific area by ID.
// @Tags city_planning
// @Produce json
// @Security BearerAuth
// @Param areaId path string true "Area ID"
// @Success 200 {object} city_planning.DemographicDataList
// @Failure 500 {object} string "Internal Server Error"
// @Router /planning/demographics/{areaId} [get]
func (h *CityPlanningHandler) GetPlanningProposalsDemographicsAria(c *gin.Context) {
	id := c.Param("areaId")
	req := &city_planning.DemographicDataFilter{ZoneId: id}

	resp, err := h.DemographicData.GetDemographicDatas(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetPlanningProposalsInfrastructure handles GET /planning/infrastructure/:infrastructureId/status
// @Summary Get infrastructure status
// @Description Get status of a specific infrastructure by ID.
// @Tags city_planning
// @Produce json
// @Security BearerAuth
// @Param infrastructureId path string true "Infrastructure ID"
// @Success 200 {object} city_planning.InfrastuctureAssetList
// @Failure 500 {object} string "Internal Server Error"
// @Router /planning/infrastructure/{infrastructureId}/status [get]
func (h *CityPlanningHandler) GetPlanningProposalsInfrastructure(c *gin.Context) {
	id := c.Param("infrastructureId")
	req := &city_planning.InfrastuctureAssetFilter{AssetId: id}

	resp, err := h.InfrastuctureAsset.GetInfrastructureAssets(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetPlanningProposalsPermitsRequirements handles GET /planning/permits/:permitTypeId/requirements
// @Summary Get permit requirements
// @Description Get requirements for a specific permit type by ID.
// @Tags city_planning
// @Produce json
// @Security BearerAuth
// @Param permitTypeId path string true "Permit Type ID"
// @Success 200 {object} city_planning.BuildingPermitFilter
// @Failure 500 {object} string "Internal Server Error"
// @Router /planning/permits/{permitTypeId}/requirements [get]
func (h *CityPlanningHandler) GetPlanningProposalsPermitsRequirements(c *gin.Context) {
	id := c.Param("permitTypeId")
	req := &city_planning.BuildingPermitFilter{PermitId: id}

	resp, err := h.BuildingPermit.GetBuildingPermits(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
