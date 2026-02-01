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

			authGroup.GET("/google", handlers.GoogleOAuth)
		}

		campaignGroup := v1.Group("/campaigns")
		{
			campaignGroup.POST("/", handlers.CreateCampaign)

			campaignGroup.POST("/fund", handlers.FundCampaign)

			campaignGroup.GET("/:slug", handlers.GetCampaign)
		}

		userGroup := v1.Group("/users/:userSlug")
		{
			userGroup.GET("/campaigns", handlers.GetUserCampaigns)
		}

		paymentGroup := v1.Group("/payments")
		{
			paymentGroup.POST("/paystack-webhook", handlers.PaystackWebhook)
			paymentGroup.POST("/initialize", handlers.InitializePayment)
		}
	}
}
