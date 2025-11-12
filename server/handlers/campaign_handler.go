package handlers

import (
	"log"
	"net/http"

	"github.com/donny-c-1/sorcemoola/server/database"
	"github.com/donny-c-1/sorcemoola/server/models"
	"github.com/gin-gonic/gin"
)

func CreateCampaign(c *gin.Context) {
	var json struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// todo Sanitize json data

	newCampaign := models.Campaign{
		Name:        json.Name,
		Description: json.Description,
	}

	result := database.DB.Create(&newCampaign)

	if result.Error != nil {
		log.Printf("Failed to create campaign: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to created campaign. Please try again.",
		})
		return
	}

	log.Printf("New campaign created with ID: %d\n", newCampaign.ID)

	c.JSON(http.StatusOK, gin.H{
		"message":  "Campaign created",
		"campaign": newCampaign,
	})
}
