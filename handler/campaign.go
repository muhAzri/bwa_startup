package handler

import (
	"bwa_startup/campaign"
	"bwa_startup/helper"
	"bwa_startup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service: service}
}

// API endpoint: GET /api/v1/campaigns
func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userIDStr := c.Query("user_id")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	var page, limit *int

	var userID *string
	if userIDStr != "" {
		userID = &userIDStr
	}

	if pageStr != "" {
		pageInt, err := strconv.Atoi(pageStr)
		if err != nil {
			response := helper.ApiResponse("Invalid page parameter", http.StatusBadRequest, "error", nil, err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}
		page = &pageInt
	}

	if limitStr != "" {
		limitInt, err := strconv.Atoi(limitStr)
		if err != nil {
			response := helper.ApiResponse("Invalid limit parameter", http.StatusBadRequest, "error", nil, err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}
		limit = &limitInt
	}

	campaigns, err := h.service.GetCampaign(userID, page, limit)

	if err != nil {
		response := helper.ApiResponse("Failed to fetch campaigns", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Campaigns retrieved successfully", http.StatusOK, "success", campaign.FormatCampaigns(campaigns), nil)
	c.JSON(http.StatusOK, response)
}

// API endpoint: GET /api/v1/campaign/:id
func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)

	if err != nil {
		response := helper.ApiResponse("Failed to get campaign details: Invalid request", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign details: Campaign not found", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Campaign details retrieved successfully", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail), nil)
	c.JSON(http.StatusOK, response)
}

// API endpoint POST /api/v1/campaign
func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		response := helper.ApiResponse("Failed to create campaign: Invalid request", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)

	if err != nil {
		response := helper.ApiResponse("Failed to create campaign", http.StatusBadRequest, "error", nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Campaigns created successfully", http.StatusOK, "success", campaign.FormatCampaign(newCampaign), nil)
	c.JSON(http.StatusOK, response)
}
