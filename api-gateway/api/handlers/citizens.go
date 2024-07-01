package handlers

import (
	"context"
	"net/http"

	"gateway/genproto/citizen"

	"github.com/gin-gonic/gin"
)


// @Summary Create a new citizen
// @Description Register a new citizen with the provided details.
// @Tags citizens
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param citizen body citizen.CitizenCreate true "Citizen Data"
// @Success 201 {object} citizen.Citizen
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /citizens/create [POST]
func (h *CitizenHandler) CreateCitizen(c *gin.Context) {
    var req citizen.CitizenCreate
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    resp, err := h.Citizen.RegisterCitizen(context.Background(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, resp)
}



// @Summary Create a new service request
// @Description Create a new service request for a citizen.
// @Tags citizens
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param citizenId path string true "Citizen ID"
// @Param request body citizen.ServiceReqCreate true "Service Request Data"
// @Success 201 {object} citizen.ServiceReq
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /citizens/{citizenId}/requests [post]
func (h *CitizenHandler) CreateCitizensRequests(c *gin.Context) {
    var req citizen.ServiceReqCreate
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    resp, err := h.CitizenRequest.CreateServiceRequest(context.Background(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, resp)
}


// @Summary Create documents for a citizen
// @Description Create documents for a specific citizen.
// @Tags citizens
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param citizenId path string true "Citizen ID"
// @Param document body citizen.ById true "Document Data"
// @Success 201 {object} citizen.DocumentList
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /citizens/{citizenId}/documents [post]
func (h *CitizenHandler) CreateCitizensDocuments(c *gin.Context) {
    var req citizen.ById
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    resp, err := h.CitizenDocument.GetCitizenDocuments(context.Background(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, resp)
}


// @Summary Get citizen profile
// @Description Get the profile details of a citizen by ID.
// @Tags citizens
// @Produce json
// @Security BearerAuth
// @Param citizenId path string true "Citizen ID"
// @Success 200 {object} citizen.Citizen
// @Failure 500 {object} string "Internal Server Error"
// @Router /citizens/{citizenId}/getbyid [get]
func (h *CitizenHandler) GetCitizen(c *gin.Context) {
    id := c.Param("citizenId")
    req := &citizen.ById{Id: id}

    resp, err := h.Citizen.GetCitizenProfile(context.Background(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}


// @Summary Get citizen requests
// @Description Get all service requests made by a specific citizen.
// @Tags citizens
// @Produce json
// @Security BearerAuth
// @Param citizenId path string true "Citizen ID"
// @Success 200 {object} citizen.ServiceReqList
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /citizens/{citizenId}/requests [get]
func (h *CitizenHandler) GetCitizensrequests(c *gin.Context) {
    id := c.Param("citizenId")
    req := &citizen.ById{Id: id}

    resp, err := h.CitizenRequest.GetCitizenRequests(context.Background(), req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}


// @Summary Update citizen profile
// @Description Update the profile details of a citizen by ID.
// @Tags citizens
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param citizenId path string true "Citizen ID"
// @Param citizen body citizen.CitizenCreate true "Updated Citizen Data"
// @Success 200 {object} citizen.Citizen
// @Failure 400 {object} string "Bad Request"
// @Router /citizens/{citizenId}/update [put]
func (h *CitizenHandler) UpdateCitizen(c *gin.Context) {
    id := c.Param("citizenId")
    var req citizen.CitizenCreate
    req.CitizenId = id

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    resp, err := h.Citizen.UpdateCitizenProfile(context.Background(), &req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resp)
}
