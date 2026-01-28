package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/donny-c-1/sorcemoola/server/database"
	"github.com/donny-c-1/sorcemoola/server/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateCampaign(c *gin.Context) {
	var json struct {
		CampaignName  string `json:"campaignName" binding:"required"`
		CampaignStory string `json:"campaignStory" binding:"required"`
		CreatedBy     string `json:"createdBy" binding:"required"`
		FundGoal      int64  `json:"fundGoal" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input" + err.Error(),
		})
		return
	}

	// todo Sanitize json data

	newCampaign := models.Campaign{
		Name:      json.CampaignName,
		Story:     json.CampaignStory,
		CreatedBy: json.CreatedBy,
		FundGoal:  json.FundGoal,
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
		"id": newCampaign.ID,
	})
}

func GetCampaign(c *gin.Context) {
	slug := c.Param("slug")

	var campaign models.Campaign
	result := database.DB.Preload("Creator", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Preload("RecentContributions", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC").Limit(3).Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		})
	}).Where("slug = ?", slug).First(&campaign)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid campaign slug",
		})
		return
	}

	c.JSON(http.StatusOK, campaign)
	// 	gin.H{
	// 	"name":  campaign.Name,
	// 	"story": campaign.Story,
	// 	"slug":  campaign.Slug,
	// 	"creator": gin.H{
	// 		"id":   campaign.Creator.ID,
	// 		"name": campaign.Creator.Name,
	// 	},
	// })
}

func GetUserCampaigns(c *gin.Context) {
	userSlug := c.Param("userSlug")

	userUUID, err := uuid.Parse(userSlug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID format",
		})
		return
	}

	var campaigns []models.Campaign
	result := database.DB.Where("created_by = ?", userUUID).Find(&campaigns)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch campaigns",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"campaigns": campaigns,
	})
}

func FundCampaign(c *gin.Context) {
	var json struct {
		CampaignID string `json:"campaignID"`
		UserID     string `json:"userID"`
		Amount     int64  `json:"amount"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err, "message": "Invalid body format"})
		return
	}
	fmt.Printf("%#v\n", json)

	userUUID, err := uuid.Parse(json.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID format"})
		return
	}

	campaignUUID, err := uuid.Parse(json.CampaignID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Campaign ID format"})
		return
	}

	log.Print("Got here")

	//* Start the transaction
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		// * 1. Create the Contribution
		contribution := models.Contribution{
			CampaignID: campaignUUID,
			UserID:     userUUID,
			Amount:     json.Amount,
		}

		if err := tx.Create(&contribution).Error; err != nil {
			return err
		}

		return tx.Model(&models.Campaign{}).Where("id = ?", campaignUUID).Updates(map[string]interface{}{
			"amount_raised": gorm.Expr("amount_raised + ?", json.Amount),
			"backers_count": gorm.Expr("backers_count + ?", 1),
		}).Error
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process contribution"})
		return
	}

	// jsonData, err := json.Marshal(paramsMap)
	// if err != nil {
	// 	fmt.Println("Error marshalling JSON: ", err)
	// 	return
	// }

	// params := string(jsonData)

	// req, err := http.NewRequest("POST", "https://api.paystack.co/transaction/initialize", bytes.NewBuffer([]byte(params)))
	// if err != nil {
	// 	fmt.Println("Error creating request: ", err)
	// 	return
	// }

	// // Set Headers
	// req.Header.Set("Authorization", "Bearer sk_test_b7511dc0e790ee7264d6118cfc0cef4e23e8cc7d")
	// req.Header.Set("Content-Type", "application/json")

	// // Create HTTP client and execute the request
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println("Error sending request: ", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// // Read the response
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response: ", err)
	// 	return
	// }

	// // Parse and print the JSON response
	// var result map[string]interface{}
	// err = json.Unmarshal(body, &result)
	// if err != nil {
	// 	fmt.Println("Error parsing response: ", err)
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"status": "succes"})
}
