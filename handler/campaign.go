package handler

import (
	"net/http"
	"strconv"

	"github.com/arif-rizal1122/bwa-startup/campaign"
	"github.com/arif-rizal1122/bwa-startup/helper"
	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	// fields here
	service campaign.Service
}


func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service: service}
}



func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	// Ambil nilai user_id dari query string
	userIDStr := c.Query("user_id")
	// Validasi apakah userIDStr adalah bilangan bulat
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID == 0 {
		response := helper.APIResponse("Invalid user_id", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Panggil service untuk mendapatkan kampanye berdasarkan userID
	campaigns, err := h.service.GetCampaings(userID)
	if err != nil {
		response := helper.APIResponse("Errors to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("shows list campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
}
