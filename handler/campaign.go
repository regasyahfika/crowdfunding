package handler

import (
	"crowdfunding/campaign"
	"crowdfunding/helper"
	"crowdfunding/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)

	if err != nil {
		response := helper.APIReseponse("Error to get campaigns", http.StatusBadRequest, "error", campaign.FormatCampaigns(campaigns))

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIReseponse("List of campaigns", http.StatusOK, "Success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIReseponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIReseponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIReseponse("Campaign detail", http.StatusOK, "Success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"erros": errors}

		response := helper.APIReseponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// get data user from authMiddleware
	currentUser := c.MustGet("currentUser").(user.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIReseponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIReseponse("Success to create campaign", http.StatusOK, "Success", campaign.FormatCampaignDetail(newCampaign))
	c.JSON(http.StatusOK, response)
	return

}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIReseponse("Failed to get detail of campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"erros": errors}

		response := helper.APIReseponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIReseponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIReseponse("Success to update campaign", http.StatusOK, "Success", campaign.FormatCampaignDetail(updatedCampaign))
	c.JSON(http.StatusOK, response)
	return
}
