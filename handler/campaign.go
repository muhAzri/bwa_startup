package handler

import (
	"bwa_startup/campaign"
	"bwa_startup/helper"
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

//api/v1/campaign

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
		response := helper.ApiResponse("Error fetching campaigns", http.StatusInternalServerError, "error", nil, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("List of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns), nil)
	c.JSON(http.StatusOK, response)

}
