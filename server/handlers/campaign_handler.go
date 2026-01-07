package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/donny-c-1/sorcemoola/server/database"
	"github.com/donny-c-1/sorcemoola/server/models"
	"github.com/gin-gonic/gin"
)

func CreateCampaign(c *gin.Context) {
	var json struct {
		CampaignName  string `json:"campaignName" binding:"required"`
		CampaignStory string `json:"campaignStory" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input" + err.Error(),
		})
		return
	}

	// todo Sanitize json data

	newCampaign := models.Campaign{
		Name:  json.CampaignName,
		Story: json.CampaignStory,
	}

	log.Print(newCampaign)
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

func FundCampaign(c *gin.Context) {
	paramsMap := map[string]string{
		"email":  "chikwemdonald@gmail.com",
		"amount": "10000",
	}

	jsonData, err := json.Marshal(paramsMap)
	if err != nil {
		fmt.Println("Error marshalling JSON: ", err)
		return
	}

	params := string(jsonData)

	req, err := http.NewRequest("POST", "https://api.paystack.co/transaction/initialize", bytes.NewBuffer([]byte(params)))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	// Set Headers
	req.Header.Set("Authorization", "Bearer sk_test_b7511dc0e790ee7264d6118cfc0cef4e23e8cc7d")
	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response: ", err)
		return
	}

	// Parse and print the JSON response
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing response: ", err)
		return
	}

	c.JSON(http.StatusOK, result)
}
