package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/donny-c-1/sorcemoola/server/database"
	"github.com/donny-c-1/sorcemoola/server/models"
	"github.com/donny-c-1/sorcemoola/server/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func InitializePayment(c *gin.Context) {
	var req struct {
		Amount     int64  `json:"amount" binding:"required"`
		UserID     string `json:"userID" binding:"required"`
		CampaignID string `json:"campaignID" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", req.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var campaign models.Campaign
	if err := database.DB.First(&campaign, "id = ?", req.CampaignID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campaign does not exist"})
		return
	}

	paystackSecret := os.Getenv("PAYSTACK_SECRET")

	body := map[string]any{
		"amount": req.Amount,
		"email":  user.Email,
		"metadata": map[string]string{
			"campaign_id": campaign.ID.String(),
			"user_id":     user.ID.String(),
		},
	}

	jsonBody, _ := json.Marshal(body)

	client := &http.Client{Timeout: 10 * time.Second}
	reqUrl := "https://api.paystack.co/transaction/initialize"

	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment request"})
		return
	}

	httpReq.Header.Set("Authorization", "Bearer "+paystackSecret)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Paystack communication error"})
		return
	}
	defer resp.Body.Close()

	var result struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Data    struct {
			AccessCode string `json:"access_code"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse payment response"})
		return
	}

	if !result.Status {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Message})
		return
	}

	c.JSON(http.StatusOK, result)
}

func PaystackWebhook(c *gin.Context) {
	ps := services.NewPaystackService()

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	signature := c.GetHeader("x-paystack-signature")
	if !ps.VerifySignature(body, signature) {
		c.Status(http.StatusUnauthorized)
		return
	}

	var event struct {
		Event string `json:"event"`
		Data  struct {
			Amount    int64  `json:"amount"`
			Reference string `json:"reference"`
			Metadata  struct {
				CampaignID string `json:"campaign_id"`
				UserID     string `json:"user_id"`
			} `json:"metadata"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &event); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	campaignUUID, err := uuid.Parse(event.Data.Metadata.CampaignID)
	if err != nil {
		log.Printf("Metadata Error: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}

	userUUID, err := uuid.Parse(event.Data.Metadata.UserID)
	if err != nil {
		log.Printf("Metadata Error: %v", err)
		c.Status(http.StatusBadRequest)
		return
	}

	if event.Event == "charge.success" {
		err := database.DB.Transaction(func(tx *gorm.DB) error {
			contribution := models.Contribution{
				ID:         uuid.New(),
				Amount:     event.Data.Amount,
				CampaignID: campaignUUID,
				UserID:     userUUID,
			}

			if err := tx.Create(&contribution).Error; err != nil {
				return err
			}

			return tx.Model(&models.Campaign{}).Where("id = ?", campaignUUID).Updates(map[string]any{
				"amount_raised": gorm.Expr("amount_raised + ?", event.Data.Amount),
				"backers_count": gorm.Expr("backers_count + ?", 1),
			}).Error
		})

		if err != nil {
			log.Printf("DATABASE ERROR: %v", err)
			c.Status(http.StatusInternalServerError)
			return
		}
	}

	c.Status(http.StatusOK)
}
