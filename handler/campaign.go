package handler

import (
	"net/http"
	"strconv"

	"github.com/arif-rizal1122/bwa-startup/campaign"
	"github.com/arif-rizal1122/bwa-startup/helper"
	"github.com/arif-rizal1122/bwa-startup/user"
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






func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaingDetailInput

	// MENGIKAT URI KE STRUCT INPUT
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to get detail of campaign", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	
	campaignDetail, err := h.service.GetCampaingByID(input)
	if err != nil {
		response := helper.APIResponse("failed to get detail of campaign", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}


	response := helper.APIResponse("succes get detail campaing", http.StatusOK, "success", campaign.FormatCampaignDetailFormatters(campaignDetail))
	c.JSON(http.StatusOK, response)
}



func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)

}




func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
   var inputID campaign.GetCampaingDetailInput

   err := c.ShouldBindUri(&inputID)
   if err != nil {
	   response := helper.APIResponse("campaign updated failed", http.StatusBadRequest, "failed", nil)
	   c.JSON(http.StatusBadRequest, response)
	   return
   }


   var inputData campaign.CreateCampaignInput
   err = c.ShouldBindJSON(&inputData)
   if err != nil {
	   errors := helper.FormatValidationError(err)
	   errorMessage := gin.H{"errors": errors}

	   response := helper.APIResponse("campaign updated failed", http.StatusUnprocessableEntity, "error", errorMessage)
	   c.JSON(http.StatusUnprocessableEntity, response)
	   return
   }
   
   currentUser := c.MustGet("currentUser").(user.User)
   inputData.User = currentUser
     
   // panggil service
   updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("campaign updated failed", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

   response := helper.APIResponse("campaign updated success", http.StatusOK, "success", campaign.FormatCampaign(updatedCampaign))
   c.JSON(http.StatusOK, response)

}