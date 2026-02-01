package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/donny-c-1/sorcemoola/server/auth"
	"github.com/donny-c-1/sorcemoola/server/database"
	"github.com/donny-c-1/sorcemoola/server/models"
	"github.com/donny-c-1/sorcemoola/server/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var json struct {
		Name        string `json:"name" binding:"required"`
		Email       string `json:"email" binding:"required"`
		Password    string `json:"password" binding:"required"`
		AccountType string `json:"accountType" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"label": "server",
			"error": "Invalid input: " + err.Error(),
		})
		return
	}

	// todo Sanitize form data

	newUser := models.User{
		Name:        json.Name,
		Email:       json.Email,
		Password:    json.Password,
		AccountType: json.AccountType,
	}

	result := database.DB.Create(&newUser)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{
				"label":   "email",
				"message": "Email already in use",
			})
			return
		}

		log.Printf("Failed to created user: %v", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"label":   "server",
			"message": "Failed to register user",
		})
		return
	}

	log.Printf("New user created with ID: %d\n", newUser.ID)

	token, err := auth.GenerateToken(&newUser, 72*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"label":   "server",
			"message": "Error Generating Token",
		})
		log.Println("Error from here too")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":          newUser.ID,
			"name":        newUser.Name,
			"email":       newUser.Email,
			"accountType": newUser.AccountType,
			"createdAt":   newUser.CreatedAt,
		},
		"token": token,
	})
}

func Login(c *gin.Context) {
	var json struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"label":   "server",
			"message": "Invalid input",
		})
		return
	}

	var user models.User
	result := database.DB.Where("email = ?", json.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"label":   "email",
			"message": "Invalid email or password",
		})
		return
	}

	if !user.VerifyPassword(json.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"label":   "password",
			"message": "Invalid email or password",
		})
		return
	}

	token, err := auth.GenerateToken(&user, 72*time.Hour)
	if err != nil {
		log.Printf("Token generation failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"label":   "server",
			"message": "Failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"id":          user.ID,
			"name":        user.Name,
			"email":       user.Email,
			"accountType": user.AccountType,
			"createdAt":   user.CreatedAt,
		},
		"token": token,
	})
}

func VerifyTokenHandler(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Missing authorization header",
		})
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid authorization format",
		})
		return
	}

	claims, err := auth.ValidateToken(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "token valid",
		"user_id": claims["user_id"],
		"email":   claims["email"],
	})
}

func GoogleOAuth(c *gin.Context) {
	code := c.Query("code")

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization code not provided!"})
		return
	}

	tokenRes, err := services.GetGoogleOauthToken(code)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	google_user, err := services.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", google_user.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// User doesn't exist, create a new one
			newUser := models.User{
				ID:          uuid.New(),
				Name:        google_user.Name,
				Email:       google_user.Email,
				AccountType: "individual",
				Password:    "",
			}
			if result := database.DB.Create(&newUser); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user"})
				return
			}
			user = newUser
		} else {
			// other database error
			c.JSON(http.StatusInternalServerError, gin.H{"message": "database error", "error": err.Error()})
			return
		}
	}

	token, err := auth.GenerateToken(&user, 72*time.Hour)
	if err != nil {
		log.Printf("Token generation failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"label":   "server",
			"message": "Failed to generate token",
		})
		return
	}

	redirectUrl := os.Getenv("FRONTEND_REDIRECT_URL")
	c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s?token=%s", redirectUrl, token))
}
