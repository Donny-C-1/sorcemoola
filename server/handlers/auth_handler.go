package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/donny-c-1/sorcemoola/server/database"
	"github.com/donny-c-1/sorcemoola/server/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var json struct {
		FirstName     string    `json:"firstname" binding:"required"`
		LastName      string    `json:"lastname" binding:"required"`
		PreferredName string    `json:"preferredname"`
		Email         string    `json:"email" binding:"required"`
		Password      string    `json:"password" binding:"required"`
		DateOfBirth   time.Time `json:"dateofbirth"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// todo Sanitize form data

	newUser := models.User{
		FirstName:     json.FirstName,
		LastName:      json.LastName,
		PreferredName: json.PreferredName,
		Email:         json.Email,
		Password:      json.Password,
		DateOfBirth:   json.DateOfBirth,
	}

	result := database.DB.Create(&newUser)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User with this email already exists",
			})
			return
		}

		log.Printf("Failed to created user: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user. Please try again.",
		})
		return
	}

	log.Printf("New user created with ID: %d\n", newUser.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    newUser,
	})
}
