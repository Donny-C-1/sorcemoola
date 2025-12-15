package routes

import (
	"github.com/donny-c-1/sorcemoola/server/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", handlers.HomeHandler)

	router.GET("/health", handlers.HealthHandler)

	v1 := router.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", handlers.Register)

			authGroup.POST("/login", handlers.Login)

			authGroup.GET("/verify", handlers.VerifyTokenHandler)
		}

		campaignGroup := v1.Group("/campaigns")
		{
			campaignGroup.POST("/", handlers.CreateCampaign)

			campaignGroup.GET("/fund", handlers.FundCampaign)
		}
	}
}
